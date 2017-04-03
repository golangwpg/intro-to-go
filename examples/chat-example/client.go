package main

import (
	"fmt"
	"net"
	"time"
)

const addr = "localhost:7070"

func main() {
	fmt.Println("Starting client...")

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello world!"))
	if err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Second)

	_, err = conn.Write([]byte("hello world again!"))
	if err != nil {
		panic(err)
	}
}
