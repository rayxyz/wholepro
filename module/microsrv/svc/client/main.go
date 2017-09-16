package main

import (
	"context"
	"fmt"
	"wholepro/module/microsrv/svc/proto"

	micro "github.com/micro/go-micro"
)

type helloClient struct{}

var client helloClient

func (client *helloClient) hello() {
	service := micro.NewService(micro.Name("ray.service.hello.client"))
	helloer := proto.NewHelloClient("ray.service.hello", service.Client())
	resp, err := helloer.SayHello(context.TODO(), &proto.HelloRequest{Name: "Xiaoming"})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.Msg)
}

func main() {
	client.hello()
}
