package main_test

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/Shresth72/sysd/ntw/unix"
	"github.com/stretchr/testify/assert"
)

func TestEchoServerUnixPacket(t *testing.T) {
	dir, err := os.MkdirTemp("", "echo_unixpacket")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if rErr := os.RemoveAll(dir); rErr != nil {
			t.Error(rErr)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))

	rAddr, err := main.StreamingEchoServer(ctx, "unixpacket", socket)
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	err = os.Chmod(socket, os.ModeSocket|0666)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := net.Dial("unixpacket", rAddr.String())
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	msg := []byte("ping")
	for i := 0; i < 3; i++ { // write 3 "ping" messages
		_, err = conn.Write(msg)
		if err != nil {
			t.Fatal(err)
		}
	}

	buf := make([]byte, 1024)
	for i := 0; i < 3; i++ { // read 3 times from the server
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, msg, buf[:n])
	}

	for i := 0; i < 3; i++ { // write 3 more "ping" messages
		_, err = conn.Write(msg)
		if err != nil {
			t.Fatal(err)
		}
	}

	buf = make([]byte, 2)    // only read the first 2 bytes of each reply
	for i := 0; i < 3; i++ { // read 3 times from the server
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotEqual(t, msg, buf[:n])
    assert.Equal(t, msg[:2], buf[:n])
	}
}
