package main
import "bufio"
import "crypto/tls"
import . "fmt"

func main() {
  if certificate, e := tls.LoadX509KeyPair("client.cert.pem", "client.key.pem"); e == nil {
    config := tls.Config{
      Certificates: []tls.Certificate{ certificate },
      InsecureSkipVerify: true,
    }

    if connection, e := tls.Dial("tcp", ":1025", &config); e == nil {
      defer connection.Close()
      if text, e := bufio.NewReader(connection).ReadString('\n'); e == nil {
        Printf(text)
      }
    }
  }
}
