package tcp_test

import (
	"crypto/rand"
	"io"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadIntoBuffer(t *testing.T) {
	// Server
	payload := make([]byte, 1<<24) // send 16Mb
	_, err := rand.Read(payload)
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)
		if err != nil {
			t.Error(err)
		}
	}()

	// Client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1<<19) // read 512Kb
	var receivedData []byte
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
			break
		}

		// t.Logf("read %d bytes", n)
		receivedData = append(receivedData, buf[:n]...)
	}

	assert.Equal(t, payload, receivedData)
}

func TestFileReadIntoBuffer(t *testing.T) {
	// Server
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		file, err := os.Open("buffer_file.txt")
		if err != nil {
			t.Error(err)
			return
		}
		defer file.Close()

		_, err = io.Copy(conn, file)
		if err != nil {
			t.Error(err)
		}
	}()

	// Read expected content from file
	expectedFile, err := os.Open("buffer_file.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer expectedFile.Close()

	expectedContent, err := io.ReadAll(expectedFile)
	if err != nil {
		t.Fatal(err)
	}

	// Client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	var receivedData []byte
	buf := make([]byte, 1<<19) // read 512Kb
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
			break
		}
		t.Logf("read %d bytes", n)
		receivedData = append(receivedData, buf[:n]...)
	}

	assert.Equal(t, expectedContent, receivedData)
}
