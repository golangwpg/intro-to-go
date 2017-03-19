package main

import "fmt"

func main() {

	// create channel (unbuffered)
	ch := make(chan (string)) // HL

	// send to channel
	go func() {
		ch <- "hey" // HL
	}()

	// receive from channel
	msg := <-ch // HL

	fmt.Printf("received: %s\n", msg)
}
