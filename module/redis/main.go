package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

func init() {
	fmt.Println("Initiating...")
}

func main() {
	fmt.Println("In redis....")
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal("Connect to redis failed.")
	}
	defer c.Close()

	obj, err := c.Do("SET", "session_id", "flfkj49857fljdflkjdsfkljdfljlkfs")
	if err != nil {
		log.Println("Redis sets value error.")
	}
	fmt.Printf("The returned interface: %#v\n", obj)
	data, errx := redis.String(c.Do("GET", "session_id"))
	if errx != nil {
		log.Println("Redis gets value error.")
	}
	fmt.Printf("The gotten value: %#v\n", data)
	fmt.Println("The gotten value in string: ", string(data))
	fmt.Println("I am dying...")
}
