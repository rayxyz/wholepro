package main

import (
	"log"
	"net"

	"golang.org/x/net/context"

	pb "wholepro/module/prototest/grpc/proto"

	"google.golang.org/grpc"
)

type CStructServer struct {
}

func (s *CStructServer) GetHiMsg(ctx context.Context, in *pb.ClientRequest) (*pb.ServerReply, error) {
	log.Println("Client request is arraving...")
	return &pb.ServerReply{Msg: "Client requested: " + in.Msg + ", I reply: Okay!!!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatal("Create gRPC server error.")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCStructServer(grpcServer, &CStructServer{})
	grpcServer.Serve(lis)
}
