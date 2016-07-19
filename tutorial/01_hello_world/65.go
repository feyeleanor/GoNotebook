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
  "reflect"
)

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

type Exception interface {
  error
}

func Raise(message string, parameters ...interface{}) {
  panic(Exception(fmt.Errorf(message, parameters...)))
}

type LaunchException error

func RaiseLaunchException(message string, parameters ...interface{}) {
  panic(LaunchException(fmt.Errorf(message, parameters...)))
}

type KeyException error

func RaiseKeyException(message string, parameters ...interface{}) {
  panic(KeyException(fmt.Errorf(message, parameters...)))
}

type CryptoException error

func RaiseCryptoException(message string, parameters ...interface{}) {
  panic(CryptoException(fmt.Errorf(message, parameters...)))
}

func Rescue(f func(), r ...interface{}) {
  defer func() {
    if e := recover(); e != nil {
      if e, ok := e.(Exception); ok {
        et := reflect.TypeOf(e)
        for _, handler := range r {
          if h := reflect.ValueOf(handler); h.Kind() == reflect.Func && h.Type().NumIn() == 1 {
            switch hpt := h.Type().In(0); {
            case et == hpt:
              fallthrough
            case hpt.Kind() == reflect.Interface && et.Implements(hpt):
              h.Call([]reflect.Value{ reflect.ValueOf(e) })
              return
            }
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
          RaiseCryptoException("unable to encrypt server response")
        } else if n, e = connection.WriteToUDP(response, c); e != nil {
          Raise("unable to write response to client: %v", c)
        }
        return
      },
      func(k KeyException) {
        log.Println("KeyException:", k.Error())
      },
      func(e Exception) {
        log.Println("Exception:", e.Error())
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
          log.Println(e.Error())
        },
      )
    }
  })
}

func Launch(address string, f func(*UDPConn) error) {
  var connection *UDPConn

  Rescue(
    func() {
      if a, e := ResolveUDPAddr("udp", address); e != nil {
        RaiseLaunchException("unable to resolve UDP address: %v", e)
      } else if connection, e = ListenUDP("udp", a); e != nil {
        RaiseLaunchException("can't open socket for listening: %v", e)
      } else if e = f(connection); e != nil {
        Raise("connection error: %v", e)
      }
    },
    func(l LaunchException) {
      log.Println("LaunchException:", l.Error())
      os.Exit(1)
    },
    func(e Exception) {
      log.Println(e.Error())
      os.Exit(1)
    },
  )
}