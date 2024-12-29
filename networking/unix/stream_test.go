package main_test

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/Shresth72/sysd/ntw/unix"
	"github.com/stretchr/testify/assert"
)

func TestEchoServerUnix(t *testing.T) {
	// temp subdirectory containing echo server's socket file
	dir, err := os.MkdirTemp("", "echo_unix")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		// cleanup
		if rErr := os.RemoveAll(dir); rErr != nil {
			t.Error(rErr)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))

	rAddr, err := main.StreamingEchoServer(ctx, "unix", socket)
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	// Change mode and make sure, read and write access
	err = os.Chmod(socket, os.ModeSocket|0666)
	if err != nil {
		t.Fatal(err)
	}

	// make connection and send test
	conn, err := net.Dial("unix", rAddr.String())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	msg := []byte("ping")
	for i := 0; i < 3; i++ { // write 3 "ping" to the server
		_, err = conn.Write(msg)
		if err != nil {
			t.Fatal(err)
		}
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf) // read once from the server
	if err != nil {
		t.Fatal(err)
	}

	expected := bytes.Repeat(msg, 3)
	assert.Equal(t, expected, buf[:n])
}
