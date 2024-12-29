package udp_test

import (
	"context"
	"fmt"
	"net"
)

func echoServerUDP(ctx context.Context, addr string) (net.Addr, error) {
  s ,err := net.ListenPacket("udp", addr)
  if err != nil {
    return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
  }

  go func() {
    go func() {
      // Blocks on context's Done channel
      <-ctx.Done()
      _ = s.Close()
    }()

    buf := make([]byte, 1024)
    for {
      n, clientAddr, err := s.ReadFrom(buf)
      if err != nil {
        return
      }

      var resp []byte
      if string(buf[:n]) == "ping" {
        resp = []byte("pong")
      } else {
        resp = buf[:n]
      }

      _, err = s.WriteTo(resp, clientAddr)
      if err != nil {
        return
      }
    }
    }()

  return s.LocalAddr(), nil
}
