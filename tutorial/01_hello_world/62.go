package main

import (
  "bytes"
  "crypto/rand"
  "crypto/rsa"
  "crypto/sha1"
  "encoding/gob"
  "fmt"
  "log"
  . "net"
)

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

type Exception interface {
  error
}

type LaunchException []interface{}

func (l LaunchException) Error() (r string) {
  if len(l) > 0 {
    r = fmt.Sprintf(l[0].(string), l[1:]...)
  }
  return
}

func RaiseLaunchException(format string, v ...interface{}) {
  panic(LaunchException(append([]interface{}{ format }, v)))
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
  defer func() {
    if e := recover(); e != nil {
      switch e := e.(type) {
      case LaunchException:
        log.Fatalln("Launch Error:", e.Error())
      case Exception:
        log.Fatalln("Exception:", e.Error())
      default:
        panic(e)
      }
    }
  }()

  Launch(address, func(connection *UDPConn) {
    for {
      defer func() {
        if e := recover(); e != nil {
          if _, ok := e.(Exception); ok {
            panic(e)
          }
          RaiseLaunchException("serve failure %v", e)
        }
      }()

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
}

func Launch(address string, f func(*UDPConn)) {
  var connection *UDPConn

  if a, e := ResolveUDPAddr("udp", address); e != nil {
    RaiseLaunchException("unable to resolve UDP address:", e.Error())
  } else if connection, e = ListenUDP("udp", a); e != nil {
    panic(LaunchException{ "can't open socket for listening:", e.Error() })
  }
  f(connection)
}