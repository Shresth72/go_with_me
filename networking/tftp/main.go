package main

import (
	"flag"
	"log"
	"os"
)

/*
  TRIVIAL FILE TRANSFER PROTOCOL

                      Client                    Server
                        |                         |
            Sending RRQ | ----------------------> | Received RRQ
                        |                         |
                        |            /----------- | Sending Block 1
       Received Block 1 | <---------/             |
                        |                         |
  Acknowledging Block 1 | ----------------------> |
                        |             /---------- | Sending Block 2
                        |            /  /-------- |
                        |           /  /          |
       Received Block 2 | <--------/  /           |
                        |            /            |
  Acknowledging Block 1 | <---------/             |
                        | ----------------------> | Sending Block 3
                        |                         |
*/

var (
	address = flag.String("a", "127.0.0.1:69", "listen address")
	payload = flag.String("p", "test.txt", "file to serve to client")
)

func main() {
	flag.Parse()

	p, err := os.ReadFile(*payload)
	if err != nil {
		log.Fatal(err)
	}

	s := Server{Payload: p}
	log.Fatal(s.ListenAndServe(*address))
}
