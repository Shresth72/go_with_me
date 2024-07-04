package tcp_test

import (
	"bytes"
	"io"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Monitor struct {
	*log.Logger
}

// Implements io.Writer
func (m *Monitor) Write(p []byte) (int, error) {
	err := m.Output(2, string(p))
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func TestMonitor(t *testing.T) {
	var logBuffer bytes.Buffer
	// Should write to os.StdOut for non tests
	monitor := &Monitor{Logger: log.New(&logBuffer, "monitor:", 0)}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
		monitor.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		b := make([]byte, 1024)
		r := io.TeeReader(conn, monitor)

		n, err := r.Read(b)
		if err != nil && err != io.EOF {
			monitor.Println(err)
			t.Error(err)
			return
		}

		w := io.MultiWriter(conn, monitor)
		_, err = w.Write(b[:n])
		if err != nil && err != io.EOF {
			monitor.Println(err)
			t.Error(err)
			return
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
		monitor.Fatal(err)
	}

	message := "Test123\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		t.Fatal(err)
		monitor.Fatal(err)
	}

	_ = conn.Close()
	<-done

	logOutput := logBuffer.String()
	assert.Contains(t, logOutput, message)
}
