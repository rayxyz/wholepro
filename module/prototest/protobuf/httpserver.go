package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	pb "protoserver/proto"

	"github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientMsg := pb.Client{}
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		if err := proto.Unmarshal(data, &clientMsg); err != nil {
			fmt.Println(err)
		}
		// fmt.Println("id: ", clientMsg.Id, "name: ", clientMsg.Name)
		log.Printf("data: %#v", clientMsg)
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("Protobuf app.")
}
