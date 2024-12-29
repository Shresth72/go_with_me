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

func TestEchoServerUnixDatagram(t *testing.T) {
	dir, err := os.MkdirTemp("", "echo_unixgram")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if rErr := os.RemoveAll(dir); rErr != nil {
			t.Error(rErr)
		}
	}()

	// Server socket
	ctx, cancel := context.WithCancel(context.Background())
	sSocket := filepath.Join(dir, fmt.Sprintf("s%d.sock", os.Getpid()))

	serverAddr, err := main.DatagramEchoServer(ctx, "unixgram", sSocket)
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	err = os.Chmod(sSocket, os.ModeSocket|0622)
	if err != nil {
		t.Fatal(err)
	}

	// Client socket
	cSocket := filepath.Join(dir, fmt.Sprintf("c%d.sock", os.Getpid()))
	client, err := net.ListenPacket("unixgram", cSocket)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	err = os.Chmod(cSocket, os.ModeSocket|0622)
	if err != nil {
		t.Fatal(err)
	}

	msg := []byte("ping")
	for i := 0; i < 3; i++ { // write 3 "ping" msgs
		_, err = client.WriteTo(msg, serverAddr)
		if err != nil {
			t.Fatal(err)
		}
	}

	buf := make([]byte, 1024)
	for i := 0; i < 3; i++ { // read 3 "ping" msgs
		n, addr, err := client.ReadFrom(buf)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, serverAddr.String(), addr.String())
		assert.Equal(t, msg, buf[:n])
	}
}
