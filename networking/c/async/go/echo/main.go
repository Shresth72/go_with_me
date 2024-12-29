package main

import (
	"fmt"
	"net"
	"os"
)

const (
	PORT = 6970
)

func main() {

	address := fmt.Sprintf("localhost:%d", PORT)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Server is listening on: %d\n", PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		fmt.Println("OK")
		conn.Close()
	}
}
