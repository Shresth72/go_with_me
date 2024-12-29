package http_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
  First string
  Last string
}

func handlePostUser(t *testing.T) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    defer func(r io.ReadCloser) {
      _, _ = io.Copy(io.Discard, r)
      _ = r.Close()
    }(r.Body)

    if r.Method != http.MethodPost {
      http.Error(w, "", http.StatusMethodNotAllowed)
      return
    }

    var u User
    err := json.NewDecoder(r.Body).Decode(&u)
    if err != nil {
      t.Error(err)
      http.Error(w, "Decode Failed", http.StatusBadRequest)
      return
    }

    w.WriteHeader(http.StatusAccepted)
  }
}

func TestPostUser(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(handlePostUser(t)))
  defer ts.Close()

  resp, err := http.Get(ts.URL)
  if err != nil {
    t.Fatal(err)
  }
  assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)

  buf := new(bytes.Buffer)
  u := User {
    First: "Dave",
    Last: "Mustaine",
  }
  err = json.NewEncoder(buf).Encode(&u)
  if err != nil {
    t.Fatal(err)
  }

  resp, err = http.Post(ts.URL, "application/json", buf)
  if err != nil {
    t.Fatal(err)
  }
  assert.Equal(t, http.StatusAccepted, resp.StatusCode)

  _ = resp.Body.Close()
}
