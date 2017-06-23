package zipkin

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/zipkin/_thrift/gen-go/scribe"
)

// Collector represents a Zipkin trace collector, which is probably a set of
// remote endpoints.
type Collector interface {
	Collect(*Span) error
}

// ScribeCollector implements Collector by forwarding spans to a Scribe
// service, in batches.
type ScribeCollector struct {
	client        scribe.Scribe
	factory       func() (scribe.Scribe, error)
	spanc         chan *Span
	sendc         chan struct{}
	quitc         chan chan struct{}
	batch         []*scribe.LogEntry
	nextSend      time.Time
	batchInterval time.Duration
	batchSize     int
}

// NewScribeCollector returns a new Scribe-backed Collector, ready for use.
func NewScribeCollector(addr string, timeout time.Duration, batchSize int, batchInterval time.Duration) (Collector, error) {
	factory := scribeClientFactory(addr, timeout)
	client, err := factory()
	if err != nil {
		return nil, err
	}
	c := &ScribeCollector{
		client:        client,
		factory:       factory,
		spanc:         make(chan *Span),
		sendc:         make(chan struct{}),
		batch:         []*scribe.LogEntry{},
		nextSend:      time.Now().Add(batchInterval),
		batchInterval: batchInterval,
		batchSize:     batchSize,
	}
	go c.loop()
	return c, nil
}

// Collect implements Collector.
func (c *ScribeCollector) Collect(s *Span) error {
	c.spanc <- s
	return nil // accepted
}

func (c *ScribeCollector) loop() {
	tickc := time.Tick(c.batchInterval / 10)

	for {
		select {
		case span := <-c.spanc:
			c.batch = append(c.batch, &scribe.LogEntry{
				Category: "zipkin", // TODO parameterize?
				Message:  serialize(span),
			})
			if len(c.batch) >= c.batchSize {
				go c.sendNow()
			}

		case <-tickc:
			if time.Now().After(c.nextSend) {
				go c.sendNow()
			}

		case <-c.sendc:
			c.nextSend = time.Now().Add(c.batchInterval)
			if err := c.send(c.batch); err != nil {
				log.DefaultLogger.Log("err", err.Error())
				continue
			}
			c.batch = c.batch[:0]
		}
	}
}

func (c *ScribeCollector) sendNow() {
	c.sendc <- struct{}{}
}

func (c *ScribeCollector) send(batch []*scribe.LogEntry) error {
	if c.client == nil {
		var err error
		if c.client, err = c.factory(); err != nil {
			return fmt.Errorf("during reconnect: %v", err)
		}
	}
	if rc, err := c.client.Log(c.batch); err != nil {
		c.client = nil
		return fmt.Errorf("during Log: %v", err)
	} else if rc != scribe.ResultCode_OK {
		// probably transient error; don't reset client
		return fmt.Errorf("remote returned %s", rc)
	}
	return nil
}

func scribeClientFactory(addr string, timeout time.Duration) func() (scribe.Scribe, error) {
	return func() (scribe.Scribe, error) {
		a, err := net.ResolveTCPAddr("tcp", addr)
		if err != nil {
			return nil, err
		}
		socket := thrift.NewTSocketFromAddrTimeout(a, timeout)
		transport := thrift.NewTFramedTransport(socket)
		if err := transport.Open(); err != nil {
			socket.Close()
			return nil, err
		}
		proto := thrift.NewTBinaryProtocolTransport(transport)
		client := scribe.NewScribeClientProtocol(transport, proto, proto)
		return client, nil
	}
}

func serialize(s *Span) string {
	t := thrift.NewTMemoryBuffer()
	p := thrift.NewTBinaryProtocolTransport(t)
	if err := s.Encode().Write(p); err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(t.Buffer.Bytes())
}

// NopCollector implements Collector but performs no work.
type NopCollector struct{}

// Collect implements Collector.
func (NopCollector) Collect(*Span) error { return nil }

// MultiCollector implements Collector by sending spans to all collectors.
type MultiCollector []Collector

// Collect implements Collector.
func (c MultiCollector) Collect(s *Span) error {
	errs := []string{}
	for _, collector := range c {
		if err := collector.Collect(s); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}
