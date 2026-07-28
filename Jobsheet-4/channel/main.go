package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println("Receive: ", v)
	}

	msg := make(chan string, 1)
	msg <- "hello"

	select {
	case v := <-msg:
		fmt.Println("Got message: ", v)
	default:
		fmt.Println("No message")
	}
}
