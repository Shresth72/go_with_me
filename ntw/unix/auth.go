package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func handleClient(conn *net.UnixConn, groups map[string]struct{}) {
  defer conn.Close()

  if !Allowed(conn, groups) {
    _, err := conn.Write([]byte("Access Denied...\n"))
    if err != nil {
      log.Println(err)
    }
    return
  }

  _, err := conn.Write([]byte("Welcome\n"))
  if err != nil {
    log.Println(err)
    return
  }

  buffer := make([]byte, 1024)
  for {
    n, err := conn.Read(buffer)
    if err != nil {
      log.Println(err)
      return
    }

    message := string(buffer[:n])
    println("received msg:", message)

    if message == "ping\n" {
      _, err := conn.Write([]byte("pong\n"))
      if err != nil {
        log.Println(err)
        return
      }
    } else {
      _, err := conn.Write([]byte("unknown command\n"))
      if err != nil {
        log.Println(err)
        return
      }
    }
  }
}

// Requesting Peer Credentials
func Allowed(conn *net.UnixConn, groups map[string]struct{}) bool {
  if conn == nil || groups == nil || len(groups) == 0 {
    return false
  }

  // underlying file object that represents Unix domain socket connection
  file, _ := conn.File()
  defer func() { _ = file.Close() }()

  var (
    err error
    ucred *unix.Ucred
  )

  for {
    // Passing File object's descriptor with protocol-level options
    // unix.SOL_SOCKET -> socket-level option
    // unix.SO_PEERCRED -> constantly tells Linux Kernel to get Peer Credentials option
    ucred, err = unix.GetsockoptUcred(int(file.Fd()), unix.SOL_SOCKET, unix.SO_PEERCRED)

    if err == unix.EINTR {
      continue // syscall interrupted, try again
    }
    if err != nil {
      log.Println(err)
      return false
    }
    break
  }

  u, err := user.LookupId(fmt.Sprint(ucred.Uid))
  if err != nil {
    log.Println(err)
    return false
  }

  gids, err := u.GroupIds()
  if err != nil {
    log.Println(err)
    return false
  }

  for _, gid := range gids {
    if _, ok := groups[gid]; ok {
      println("valid gid", gid)
      return true
    }
  }
  return false
}

// Service
// Accepts group names found in the Linux OS's /etc/group as cmd line args
// and listens to a Unix domain socket file.
// Allows client to connect only if they are a member of any group specified
// and retrieve peer credentials of the client if authorized
func init() {
  flag.Usage = func() {
    _, _ = fmt.Fprintf(flag.CommandLine.Output(),
      "Usage:\n\t%s <group names>\n", filepath.Base(os.Args[0]))
    flag.PrintDefaults()
  }
}

func parseGroupNames(args []string) map[string]struct{} {
  groups := make(map[string]struct{})

  for _, arg := range args {
    grp, err := user.LookupGroup(arg)
    if err != nil {
      log.Println(err)
      continue
    }

    groups[grp.Gid] = struct{}{}
  }

  return groups
}
