# gokit [![Circle CI](https://circleci.com/gh/go-kit/kit.svg?style=svg)](https://circleci.com/gh/go-kit/kit) [![Drone.io](https://drone.io/github.com/go-kit/kit/status.png)](https://drone.io/github.com/go-kit/kit/latest) [![Travis CI](https://travis-ci.org/go-kit/kit.svg?branch=master)](https://travis-ci.org/go-kit/kit) [![GoDoc](https://godoc.org/github.com/go-kit/kit?status.svg)](https://godoc.org/github.com/go-kit/kit)

**Gokit** is a **distributed programming toolkit** for microservices in the modern enterprise.

- Mailing list: [go-kit](https://groups.google.com/forum/#!forum/go-kit)
- IRC: [Freenode](https://freenode.net) `#gokit`

## Motivation

See [the motivating blog post](http://peter.bourgon.org/go-kit) and [the video of the talk](https://www.youtube.com/watch?v=iFR_7AKkJFU).

## Goals

- Operate in a heterogeneous SOA — expect to interact with mostly non-gokit services
- RPC as the messaging pattern
- Pluggable serialization and transport — not just JSON over HTTP
- Zipkin-compatible request tracing

## Non-goals

- Supporting messaging patterns other than RPC — pub/sub, CQRS, etc.
- Having opinions on deployment, orchestration, process supervision, etc.
- Having opinions on configuration passing, i.e. flags, env vars, files, etc.

## Component status

- API stability — adopted
- `package metrics` — [implemented](https://github.com/go-kit/kit/tree/master/metrics)
- `package server` — [implemented](https://github.com/go-kit/kit/tree/master/server)
- `package transport` — [implemented](https://github.com/go-kit/kit/tree/master/transport)
- `package log` — [implemented](https://github.com/go-kit/kit/tree/master/log)
- `package tracing` — [prototyping](Https://github.com/go-kit/kit/tree/master/tracing)
- `package client` — pending
- Service discovery — pending

## Contributing

At this stage, we're still developing the initial drafts of all of the packages, using an
[RFC workflow](https://github.com/go-kit/kit/tree/master/rfc).
Before submitting major changes, please write to
 [the mailing list](https://groups.google.com/forum/#!forum/go-kit)
to register your interest, and check the
 [open issues](https://github.com/go-kit/kit/issues) and
 [pull requests](https://github.com/go-kit/kit/pulls)
for existing discussions.

### Dependency management

Users who import gokit into their `package main` are responsible to organize
and maintain all of their dependencies to ensure code compatibility and build
reproducibility. Gokit makes no direct use of dependency management tools like
[Godep](https://github.com/tools/godep).

We will use a variety of continuous integration providers to find and fix
compatibility problems as soon as they occur.

### API stability policy

The gokit project depends on code maintained by others. This includes the Go
standard library and sub-repositories and other external libraries.
The Go language and standard library provide stability guarantees, but the other external libraries typically do not.
[The API Stability RFC](https://github.com/go-kit/kit/tree/master/rfc/rfc007-api-stability.md)
proposes a standard policy for package authors to advertise API stability.
The go-kit project prefers to depend on code that abides the API stability policy.

## Related projects

Projects with a ★ have had particular influence on gokit's design.

### Service frameworks

- [go-micro](https://github.com/asim/go-micro), a microservices client/server library ★
- [gocircuit](https://github.com/gocircuit/circuit), dynamic cloud orchestration
- [gotalk](https://github.com/rsms/gotalk), async peer communication protocol &amp; library
- [Kite](https://github.com/koding/kite), a micro-service framework

### Individual components

- [afex/hystrix-go](https://github.com/afex/hystrix-go), client-side latency and fault tolerance library
- [armon/go-metrics](https://github.com/armon/go-metrics), library for exporting performance and runtime metrics to external metrics systems
- [codahale/lunk](https://github.com/codahale/lunk), structured logging in the style of Google's Dapper or Twitter's Zipkin
- [eapache/go-resiliency](https://github.com/eapache/go-resiliency), resiliency patterns
- [FogCreek/logging](https://github.com/FogCreek/logging), a tagged style of logging
- [grpc/grpc-go](https://github.com/grpc/grpc-go), HTTP/2 based RPC
- [inconshreveable/log15](https://github.com/inconshreveable/log15), simple, powerful logging for Go ★
- [mailgun/vulcand](https://github.com/mailgun/vulcand), prorammatic load balancer backed by etcd
- [mattheath/phosphor](https://github.com/mattheath/phosphor), distributed system tracing
- [pivotal-golang/lager](https://github.com/pivotal-golang/lager), an opinionated logging library
- [rubyist/circuitbreaker](https://github.com/rubyist/circuitbreaker), circuit breaker library
- [Sirupsen/logrus](https://github.com/Sirupsen/logrus), structured, pluggable logging for Go ★
- [sourcegraph/appdash](https://github.com/sourcegraph/appdash), application tracing system based on Google's Dapper
- [spacemonkeygo/monitor](https://github.com/spacemonkeygo/monitor), data collection, monitoring, instrumentation, and Zipkin client library
- [streadway/handy](https://github.com/streadway/handy), net/http handler filters
- [vitess/rpcplus](https://godoc.org/code.google.com/p/vitess/go/rpcplus), package rpc + context.Context
- [gdamore/mangos](https://github.com/gdamore/mangos), nanomsg implementation in pure Go

### Web frameworks

- [Beego](http://beego.me/)
- [Gin](https://gin-gonic.github.io/gin/)
- [Goji](https://github.com/zenazn/goji)
- [Gorilla](http://www.gorillatoolkit.org)
- [Martini](https://github.com/go-martini/martini)
- [Negroni](https://github.com/codegangsta/negroni)
- [Revel](https://revel.github.io/)

## Additional reading

- [Architecting for the Cloud](http://fr.slideshare.net/stonse/architecting-for-the-cloud-using-netflixoss-codemash-workshop-29852233) — Netflix
- [Dapper, a Large-Scale Distributed Systems Tracing Infrastructure](http://research.google.com/pubs/pub36356.html) — Google
- [Your Server as a Function](http://monkey.org/~marius/funsrv.pdf) (PDF) — Twitter
