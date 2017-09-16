package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/consul/api"
)

func RegisterService(ag *api.Agent, service, hostname, node, protocol string, port int) error {
	rand.Seed(time.Now().UnixNano())
	sid := rand.Intn(65534)
	serviceID := service + "-" + strconv.Itoa(sid)
	address := protocol + "://" + hostname + ":" + strconv.Itoa(port)
	log.Println("address => ", address)
	consulService := api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    service,
		Address: address,
		Tags:    []string{node},
		Port:    port,
		// Check: &api.AgentServiceCheck{
		// 	Script:   "curl --connect-timeout=3 " + address,
		// 	Interval: "5s",
		// 	Timeout:  "8s",
		// 	TTL:      "",
		// 	HTTP:     protocol + "://" + hostname + ":" + strconv.Itoa(port),
		// 	Status:   "passing",
		// },
		Checks: api.AgentServiceChecks{},
	}

	err := ag.ServiceRegister(&consulService)
	if err != nil {
		return err
	}

	return err
}

func runService(serviceName string) {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Println(err)
	}
	// services, err := client.Agent().Services()
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// for _, v := range services {
	// 	fmt.Println("ServiceId => ", v.ID, " ServiceName => ", v.Service, " address => ", v.Address)
	// }

	servs, _, err := client.Catalog().Service(serviceName, "", nil)
	if err != nil {
		log.Println(err)
		return
	}

	if len(servs) > 0 {
		// log.Println(meta)
		serv := servs[0]
		log.Println("service address in srv => ", serv.ServiceAddress)
		resp, err := http.Get(serv.ServiceAddress)
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
				return
			}
		}
		fmt.Println(string(body))
	}
}

func main() {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	hostname := "localhost"

	if err != nil {
		fmt.Printf("Encountered error connecting to consul on %s => %s\n", hostname, err)
		return
	}

	if err = RegisterService(client.Agent(), "com.xyz.ray.hello", hostname, "localhost", "http", 8085); err != nil {
		fmt.Printf("Encountered error registering a service with consul -> %s\n", err)
	}

	runService("com.xyz.ray.hello")
}
