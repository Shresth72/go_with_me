package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
)

func main() {
  flag.Parse()

  groups := parseGroupNames(flag.Args())
  socket := filepath.Join(os.TempDir(), "creds.sock")

  addr, err := net.ResolveUnixAddr("unix", socket)
  if err != nil {
    log.Fatal(err)
  }

  s, err := net.ListenUnix("unix", addr)
  if err != nil {
    log.Fatal(err)
  }

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt)
  go func() {
    <-c
    _ = s.Close()
  }()

  fmt.Printf("Listening on %s ...\n", socket)

  for {
    conn, err := s.AcceptUnix()
    if err != nil {
      break
    }

    go handleClient(conn, groups)
  }
}
