package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

var consulClient = NewClient()

func NewClient() *api.Client {
	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Println(err)
	}
	return client
}

func RegisterService(service, node, address string, port int) error {
	regised, err := CheckServiceRegistered(service)
	if err != nil {
		return err
	}
	if regised {
		return errors.New("Service has been registered")
	}
	log.Println("address => ", address)
	consulService := api.AgentServiceRegistration{
		ID:      service,
		Name:    service,
		Address: address,
		Tags:    []string{node},
		Port:    port,
		Checks:  api.AgentServiceChecks{},
	}

	err = consulClient.Agent().ServiceRegister(&consulService)
	if err != nil {
		return err
	}

	return err
}

func DeregisterService(service string) error {
	err := consulClient.Agent().ServiceDeregister(service)
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetService(serviceName string) (*api.CatalogService, error) {
	servs, _, err := consulClient.Catalog().Service(serviceName, "", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if len(servs) > 0 {
		// log.Println(meta)
		return servs[0], nil
	}
	return nil, nil
}

func CheckServiceRegistered(service string) (bool, error) {
	srv, err := GetService(service)
	if err != nil {
		return false, err
	}
	if srv != nil {
		return true, nil
	}
	return false, nil
}

func main() {
	address := "http://localhost"
	port := 8085
	if err := RegisterService("com.xyz.ray.hello", "localhost", address, port); err != nil {
		fmt.Printf("Encountered error registering a service with consul -> %s\n", err)
	}
}
