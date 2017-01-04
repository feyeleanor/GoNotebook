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
  "os"
)

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

type Exception interface {
  error
}

func Raise(message string, parameters ...interface{}) {
  panic(fmt.Errorf(message, parameters...))
}

func Rescue(f func()) {
  defer func() {
    if e := recover(); e != nil {
      if e, ok := e.(Exception); ok {
        log.Println("Exception:", e.Error())
        os.Exit(1)
      } else {
        panic(e)
      }
    }
  }()

  f()
}

func main() {
  Serve(":1025", func(connection *UDPConn, c *UDPAddr, packet *bytes.Buffer) (n int) {
    var key rsa.PublicKey
    var response []byte

    Rescue(func() {
      if e := gob.NewDecoder(packet).Decode(&key); e != nil {
        Raise("unable to decode wrapper: %v", c)
      } else if response, e = rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e != nil {
        Raise("unable to encrypt server response")
      } else if n, e = connection.WriteToUDP(response, c); e != nil {
        Raise("unable to write response to client: %v", c)
      }
    })
    return
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer) int) {
  Rescue(func() {
    e := Launch(address, func(connection *UDPConn) (e error) {
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

    if e != nil {
      log.Fatalln(e.Error())
    }
  })
}

func Launch(address string, f func(*UDPConn) error) error {
  var connection *UDPConn

  if a, e := ResolveUDPAddr("udp", address); e != nil {
    Raise("unable to resolve UDP address: %v", e)
  } else if connection, e = ListenUDP("udp", a); e != nil {
    Raise("can't open socket for listening: %v", e.Error())
  }
  return f(connection)
}