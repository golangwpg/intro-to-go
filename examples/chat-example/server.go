package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const addr = "localhost:7070"

func main() {
	fmt.Println("Starting server...")

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening for connections on %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	fmt.Printf("Opened connection from %s\n", conn.RemoteAddr().String())

	buf := make([]byte, 256)

	for {
		bytesRead, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Closing connection from %s\n", conn.RemoteAddr().String())
				conn.Close()
				return
			}
		}

		if bytesRead > 0 {
			fmt.Printf("Msg from %s: %s\n", conn.RemoteAddr().String(), string(buf))
		}
	}

}
