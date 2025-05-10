package main

import (
	"fmt"
	"httpfromtcp/internal/request"
	"log"
	"net"
)

func main() {
	println("tcplistener.go")
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error listening:", err)
	}
	println("Listening on localhost:42069")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}
		defer conn.Close()
		println("Connection accepted")
		req, _ := request.RequestFromReader(conn)
		fmt.Printf(
			"Request line:\n- Method: %s\n- Target: %s\n- Version: %s\n",
			req.RequestLine.Method,
			req.RequestLine.RequestTarget,
			req.RequestLine.HttpVersion,
		)
		if len(req.Headers) > 0 {
			fmt.Println("Headers:")
		}
		for k, v := range req.Headers {
			fmt.Printf("- %s: %s\n", k, v)
		}
	}
}
