package main
import "bytes"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "encoding/gob"
import . "fmt"
import . "net"

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

func main() {
  Serve(":1025", func(connection *UDPConn, c *UDPAddr, packet *bytes.Buffer) (n int) {
    var key rsa.PublicKey
    if e := gob.NewDecoder(packet).Decode(&key); e == nil {
      if response, e := rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e == nil {
        n, _ = connection.WriteToUDP(response, c)
      }
    }
    return
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer) int) {
  Launch(address, func(connection *UDPConn) {
    for {
      buffer := make([]byte, 1024)
      if n, client, e := connection.ReadFromUDP(buffer); e == nil {
        go func(c *UDPAddr, b []byte) {
          if n := f(connection, c, bytes.NewBuffer(b)); n != 0 {
            Println(n, "bytes written to", c)
          }
        }(client, buffer[:n])
      }
    }
  })
}

func Launch(address string, f func(*UDPConn)) {
  if a, e := ResolveUDPAddr("udp", address); e == nil {
    if server, e := ListenUDP("udp", a); e == nil {
      f(server)
    }
  }
}
