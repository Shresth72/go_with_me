package udp_test

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoServerUDP(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
  if err != nil {
    t.Fatal(err)
  }
  defer cancel()

  client, err := net.ListenPacket("udp", "127.0.0.1:")
  if err != nil {
    t.Fatal(err)
  }
  defer func() { _ = client.Close() }()

  msg := []byte("ping")
  _, err = client.WriteTo(msg, serverAddr)
  if err != nil {
    t.Fatal(err)
  }

  buf := make([]byte, 1024)
  n, addr, err := client.ReadFrom(buf)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, serverAddr, addr)
  assert.Equal(t, []byte("pong"), buf[:n])
}

func TestListenPacketUDP(t *testing.T) {
  ctx, cancel := context.WithCancel(context.Background())
  serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
  if err != nil {
    t.Fatal(err)
  }
  defer cancel()

  client, err := net.ListenPacket("udp", "127.0.0.1:")
  if err != nil {
    t.Fatal(err)
  }
  defer func() { _ = client.Close()}()

  // Adding an interloper and interrupting the client with a message
  interloper, err := net.ListenPacket("udp", "127.0.0.1:")
  if err != nil {
    t.Fatal(err)
  }

  interrupt := []byte("interrupting")
  n, err := interloper.WriteTo(interrupt, client.LocalAddr())
  if err != nil {
    t.Fatal(err)
  }
  _ = interloper.Close()

  if l := len(interrupt); l != n {
    t.Fatalf("wrote %d bytes of %d", n, l)
  }

  // ping
  ping := []byte("ping")
  _, err = client.WriteTo(ping, serverAddr)
  if err != nil {
    t.Fatal(err)
  }

  buf := make([]byte, 1024)
  n, addr, err := client.ReadFrom(buf)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, interrupt, buf[:n])
  assert.Equal(t, interloper.LocalAddr().String(), addr.String())

  n, addr, err = client.ReadFrom(buf)
  if err != nil {
    t.Fatal(err)
  }

  assert.Equal(t, []byte("pong"), buf[:n])
  assert.Equal(t, serverAddr.String(), addr.String())
}
