package main
import "bytes"
import "crypto/rand"
import "crypto/rsa"
import "crypto/sha1"
import "encoding/gob"
import "fmt"
import "log"
import . "net"
import "os"
import "reflect"

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

type Exception interface {
  error
}

func Raise(message string, parameters ...interface{}) {
  panic(Exception(fmt.Errorf(message, parameters...)))
}

type KeyException error

func RaiseKeyException(message string, parameters ...interface{}) {
  panic(KeyException(fmt.Errorf(message, parameters...)))
}

func attemptCall(e Exception, handler interface{}) (ok bool) {
  if h := reflect.ValueOf(handler); h.Kind() == reflect.Func {
    et := reflect.TypeOf(e)
    if hpt := h.Type().In(0); et == hpt || et.Implements(hpt) {
      h.Call([]reflect.Value{reflect.ValueOf(e)})
      return true
    }
  }
  return
}

func Rescue(f func(), r ...interface{}) {
  defer func() {
    if e := recover(); e != nil {
      if e, ok := e.(Exception); ok {
        for _, h := range r {
          if attemptCall(e, h) {
            return
          }
        }
      }
      panic(e)
    }
  }()

  f()
}

func main() {
  Serve(":1025", func(connection *UDPConn, c *UDPAddr, packet *bytes.Buffer) (n int) {
    var key rsa.PublicKey
    var response []byte

    Rescue(
      func() {
        if e := gob.NewDecoder(packet).Decode(&key); e != nil {
          RaiseKeyException("unable to decode wrapper: %v", c)
        } else if response, e = rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e != nil {
          Raise("unable to encrypt server response")
        } else if n, e = connection.WriteToUDP(response, c); e != nil {
          Raise("unable to write response to client: %v", c)
        }
        return
      },
      func(k KeyException) {
        log.Println("KeyException:", k.Error())
      },
    )
    return
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer) int) {
  Launch(address, func(connection *UDPConn) (e error) {
    for {
      Rescue(
        func() {
          buffer := make([]byte, 1024)
          if n, client, e := connection.ReadFromUDP(buffer); e == nil {
            go func(c *UDPAddr, b []byte) {
              if n := f(connection, c, bytes.NewBuffer(b)); n != 0 {
                log.Println(n, "bytes written to", c)
              }
            }(client, buffer[:n])
          } else {
            Raise("%v: %v", address, e.Error())
          }
        },
        func(e Exception) {
          log.Println("Connection Error:", e.Error())
        },
      )
    }
  })
}

func Launch(address string, f func(*UDPConn) error) {
  var connection *UDPConn

  Rescue(func() {
    if a, e := ResolveUDPAddr("udp", address); e != nil {
      Raise("unable to resolve UDP address: %v", e)
    } else if connection, e = ListenUDP("udp", a); e != nil {
      Raise("can't open socket for listening: %v", e)
    } else if e = f(connection); e != nil {
      Raise("connection error: %v", e)
    }
  })
}
