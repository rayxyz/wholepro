package main

import (
	"context"
	"fmt"
	"wholepro/module/microsrv/svc/proto"

	micro "github.com/micro/go-micro"
)

func hello() {
	service := micro.NewService()
	service.Init()
	helloer := proto.NewHelloClient("ray.service.hello", service.Client())
	resp, err := helloer.SayHello(context.TODO(), &proto.HelloRequest{Name: "Ray"}, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.Msg)
}

func main() {
	hello()
}
