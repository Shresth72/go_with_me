package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	PORT        = 6970
	BUFFER_SIZE = 1024
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

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, BUFFER_SIZE)

	n, err := bufio.NewReader(conn).Read(buffer)
	if err != nil {
		fmt.Printf("Error reading from client: %v\n", err)
		return
	}

	_, err = conn.Write(buffer[:n])
	if err != nil {
		fmt.Printf("Error writing to client: %v\n", err)
		return
	}

	// fmt.Printf("Client sent: %s\n", string(buffer[:n]))
}
