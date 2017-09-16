package main

import (
	"log"

	"golang.org/x/net/context"

	"wholepro/module/microsrv/svc/proto"

	"github.com/micro/go-micro"
)

type helloServer struct{}

func (h *helloServer) SayHello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) error {
	resp.Msg = "Hello " + req.Name + "!"
	return nil
}

func main() {
	service := micro.NewService(micro.Name("ray.service.hello"), micro.Version("0.1"))
	service.Init()
	proto.RegisterHelloHandler(service.Server(), new(helloServer))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
