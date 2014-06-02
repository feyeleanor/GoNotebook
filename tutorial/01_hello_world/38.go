package main

import (
  . "fmt"
  "net"
)

func main() {
  if listener, e := net.Listen("tcp", ":1024"); e == nil {
    for {
      if connection, e := listener.Accept(); e == nil {
        defer connection.Close()
        go func(c net.Conn) {
          Fprintln(c, "hello world")
        }(connection)
      }
    }
  }
}
