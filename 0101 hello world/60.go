package main
import "bytes"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "encoding/gob"
import "fmt"
import "log"
import . "net"

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

type LaunchError []interface{}

func (l LaunchError) Error() (r string) {
  if len(l) > 0 {
    r = fmt.Sprintf(l[0].(string), l[1:]...)
  }
  return
}

func NewLaunchError(format string, v ...interface{}) (l LaunchError) {
  return LaunchError(append([]interface{}{ format }, v))
}

func main() {
  Serve(":1025", func(connection *UDPConn, c *UDPAddr, packet *bytes.Buffer) (n int) {
    var key rsa.PublicKey
    var response []byte

    if e := gob.NewDecoder(packet).Decode(&key); e != nil {
      log.Println("unable to decode wrapper:", c)
    } else if response, e = rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e != nil {
      log.Println("unable to encrypt server response")
    } else if n, e = connection.WriteToUDP(response, c); e != nil {
      log.Println("unable to write response to client:", c)
    }
    return
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer) int) {
  e := Launch(address, func(connection *UDPConn) (e error) {
    defer func() {
      if x := recover(); x != nil {
        e = LaunchError{ "serve failure %v", x }
      }
    }()
    for {
      buffer := make([]byte, 1024)
      if n, client, e := connection.ReadFromUDP(buffer); e == nil {
        go func(c *UDPAddr, b []byte) {
          if n := f(connection, c, bytes.NewBuffer(b)); n != 0 {
            log.Println(n, "bytes written to", c)
          }
        }(client, buffer[:n])
      } else {
        log.Println(address, e.Error())
      }
    }
    return
  })

  if e, ok:= e.(LaunchError); ok {
    log.Fatalln(e.Error())
  }
}

func Launch(address string, f func(*UDPConn) error) error {
  var connection *UDPConn

  if a, e := ResolveUDPAddr("udp", address); e != nil {
    return NewLaunchError("unable to resolve UDP address:", e.Error())
  } else if connection, e = ListenUDP("udp", a); e != nil {
    return LaunchError{ "can't open socket for listening:", e.Error() }
  }
  return f(connection)
}
