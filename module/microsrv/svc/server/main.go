package main

import (
	"log"

	"golang.org/x/net/context"

	"wholepro/module/microsrv/svc/proto"

	"github.com/micro/go-micro"
)

type Helloer struct{}

func (h *Helloer) SayHello(ctx context.Context, req *proto.HelloRequest, resp *proto.HelloResponse) error {
	resp.Msg = "Hello " + req.Name + "!"
	return nil
}

// func runClient(service micro.Service) {
// 	helloer := proto.NewHelloClient("ray.service.hello", service.Client())
// 	resp, err := helloer.SayHello(context.TODO(), &proto.HelloRequest{Name: "Ray"}, nil)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println(resp.Msg)
// }

func main() {
	service := micro.NewService(micro.Name("ray.service.hello"), micro.Version("0.1"))
	service.Init(
	// micro.Action(func(ctx *cli.Context) {
	// 	if ctx.Bool("run_client") {
	// 		runClient(service)
	// 		os.Exit(0)
	// 	}
	// }),
	)
	proto.RegisterHelloHandler(service.Server(), new(Helloer))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
