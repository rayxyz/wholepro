package main

import "fmt"

func printStuff() {
	fmt.Println("Hi!")
}

func main() {
	var bigstack [1024 * 1024 * 1000]int
	p := &bigstack
	fmt.Println("Address of p: ", &p)
	printStuff()
}
