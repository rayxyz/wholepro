package main

import "fmt"

func sayMsg() {
	ch := genMsg()
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func genMsg() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	sayMsg()
}
