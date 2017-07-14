package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	// conn, err := grpc.Dial("127.0.0.1:10001")
	conn, err := grpc.Dial("127.0.0.1:21210")
	if err != nil {
		log.Fatal("Connection to server error.")
	}
	defer conn.Close()
	// c := pb.NewCStructClient(conn)
	// reply, err := c.GetHiMsg(context.Background(), &pb.ClientRequest{Msg: "Hi, server!!!"})
	if err != nil {
		log.Fatal("Get server reply error.")
	}
	// fmt.Println(reply.Msg)
	log.Println("Connected.")
}
