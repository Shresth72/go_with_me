package tcp_test

import (
	"bufio"
	"errors"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	payload = "The bigger the interface, the weaker the abstraction"
	maxTries = 7
	retryInterval = 5 * time.Second
)

// bufio.Scanner - allows reading delimited data
func TestScanner(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		var writeErr error
		for tries := maxTries; tries > 0; tries-- {
			_, writeErr = conn.Write([]byte(payload))
			if writeErr != nil {
				if netErr, ok := writeErr.(net.Error); ok && netErr.Timeout() {
					log.Println("temporary error:", netErr)
					time.Sleep(retryInterval)
					continue
				}
				t.Error(writeErr)
				break
			}
			break
		}

		if writeErr != nil {
			t.Error(errors.New("temporary write failure threshold exceeded"))
		}
	}()

	// Client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"The", "bigger", "the", "interface,", "the", "weaker", "the", "abstraction"}
	assert.Equal(t, expected, words)
}
