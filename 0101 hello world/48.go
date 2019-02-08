package main
import . "fmt"
import "net"

func main() {
  if listener, e := net.Listen("tcp", ":1024"); e == nil {
    for {
      if connection, e := listener.Accept(); e == nil {
        go func(c net.Conn) {
          defer c.Close()
          Fprintln(c, "hello world")
        }(connection)
      }
    }
  }
}
