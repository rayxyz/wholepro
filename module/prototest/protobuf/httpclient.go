package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func main() {
	myClient := Client{Id: 526, Name: "John Doe"}

	data, err := proto.Marshal(&myClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = http.Post("http://localhost:8080", "", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return
	}
}
