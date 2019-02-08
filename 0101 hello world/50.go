package main
import "crypto/rand"
import "crypto/tls"
import . "fmt"

func main() {
  if certificate, e := tls.LoadX509KeyPair("server.cert.pem", "server.key.pem"); e == nil {
    config := tls.Config{
      Certificates: []tls.Certificate{ certificate },
      Rand: rand.Reader,
    }

    if listener, e := tls.Listen("tcp", ":1025", &config); e == nil {
      for {
        if connection, e := listener.Accept(); e == nil {
          go func(c *tls.Conn) {
            defer c.Close()
            Fprintln(c, "hello world")
          }(connection.(*tls.Conn))
        }
      }
    }
  }
}
