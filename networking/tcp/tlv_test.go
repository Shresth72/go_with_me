package tcp

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayloads(t *testing.T) {
	// Server
	b1 := Binary("Clear is better than clever")
	b2 := Binary("Don't panic")
	s1 := String("Errors are values")
	payloads := []Payload{&b1, &s1, &b2}

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

		for _, p := range payloads {
			_, err = p.WriteTo(conn)
			if err != nil {
				t.Error(err)
				break
			}
		}
	}()

	// Client
	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < len(payloads); i++ {
		actual, err := decode(conn)
		if err != nil {
			t.Fatal(err)
		}

		expected := payloads[i]
		assert.Equal(t, expected, actual)
	}
}

func decode(r io.Reader) (Payload, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return nil, err
	}

	var payload Payload

	switch typ {
	case BinaryType:
		payload = new(Binary)
	case StringType:
		payload = new(String)
	default:
		return nil, errors.New("unknown type")
	}

	if _, err := payload.ReadFrom(r); err != nil {
		return nil, err
	}

	return payload, nil
}
