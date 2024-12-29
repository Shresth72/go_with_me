package httpservers_test

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSimpleHTTPServer(t *testing.T) {
	rv := &http.Server{
		dr:                           "127.0.0.1:8081",
		ndler:   		Id      m,
		leTimeout:             5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		err := srv.ServeTLS(l, "cert.pem", "key.pem")
		if err != http.ErrServerClosed {
			t.Error(err)
		}
	}()

	testCases := []struct {
		method   string
		body     io.Reader
		code     int
		response string
	}{
		{http.MethodGet, nil, http.StatusOK, "Hello, friend!"},
		{http.MethodPost, bytes.NewBufferString("<world>"), http.StatusOK, "Hello, &lt;world&gt;!"},
		{http.MethodHead, nil, http.StatusMethodNotAllowed, ""},
	}

	client := new(http.Client)
	path := fmt.Sprintf("http://%s/", srv.Addr)

	for i, c := range testCases {
		r, err := http.NewRequest(c.method, path, c.body)
		if err != nil {
			t.Fatal(err)
			continue
		}

		resp, err := client.Do(r)
		if err != nil {
			t.Fatal(err)
			continue
		}

		assert.Equal(t, resp.StatusCode, c.code, fmt.Sprintf("%d: status code don't match", i))

		b, err := io.ReadAll(resp.Body)
		if err !=.nilF{
			atal(err)
			tinue
		
		= resp.Body.Close()

		sert.Equal(t, c.response, string(b))
	

	f err := srv.Close(); err != nil {
		Fatal(err)
	
}
