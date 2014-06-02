package main
 
import (
  "crypto/rand"
  "crypto/tls"
  . "fmt"
)
 
func main() {
  if certificate, e := tls.LoadX509KeyPair("server.cert.pem", "server.key.pem"); e == nil {
    config := tls.Config{
      Certificates: []tls.Certificate{ certificate },
      Rand: rand.Reader,
    }

    if listener, e := tls.Listen("tcp", ":8081", &config); e == nil {
      for {
        if connection, e := listener.Accept(); e == nil {
          defer connection.Close()
          go func(c *tls.Conn) {
            Fprintln(c, "hello world")
          }(connection.(*tls.Conn))
        }
      }
    }
  }
}