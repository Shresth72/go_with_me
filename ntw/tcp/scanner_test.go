package tcp_test

import (
	"bufio"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// bufio.Scanner - allows reading delimited data

const payload = "The bigger the interface, the weaker the abstraction"

func TestScanner(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
  if err != nil {
    t.Fatal(err)
  }

  go func ()  {
    conn, err := listener.Accept()
    if err != nil {
      t.Error(err)
      return
    }
    defer conn.Close()

    _, err = conn.Write([]byte(payload))
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
