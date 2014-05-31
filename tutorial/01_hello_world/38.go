package main

import (
  . "fmt"
  "net"
)

func main() {
  if server, e := net.Listen("tcp", ":8080"); e == nil {
    for {
      if connection, e := server.Accept(); e == nil {
        go func(c net.Conn) {
          Fprintln(c, "hello world")
          c.Close()
        }(connection)
      }
    }
  }
}
