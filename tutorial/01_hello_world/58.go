package main

import (
  "bytes"
  "crypto/rand"
  "crypto/rsa"
  "crypto/sha1"
  "encoding/gob"
  "log"
  . "net"
)

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

func main() {
  Serve(":1025", func(connection *UDPConn, c *UDPAddr, packet *bytes.Buffer) {
    var e error
    var n int
    var key rsa.PublicKey
    var response []byte

    if e = gob.NewDecoder(packet).Decode(&key); e != nil {
      log.Println("unable to decode wrapper:", c)
    }

    if response, e = rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e != nil {
      log.Println("unable to encrypt server response")
    }

    if n, e = connection.WriteToUDP(response, c); e != nil {
      log.Println("unable to write response to client:", c)
    }
    log.Println(n, "bytes written to", c)
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer)) {
  Launch(address, func(connection *UDPConn) {
    for {
      b := make([]byte, 1024)
      if n, c, e := connection.ReadFromUDP(b); e == nil {
        go func(client *UDPAddr, buffer []byte) {
          f(connection, client, bytes.NewBuffer(buffer))
        }(c, b[:n])
      } else {
        log.Println(address, e.Error())
      }
    }
  })
}

func Launch(address string, f func(*UDPConn)) {
  var e error
  var a *UDPAddr
  var server *UDPConn

  if a, e = ResolveUDPAddr("udp", address); e != nil {
    log.Fatalln("unable to resolve UDP address:", e.Error())
  }

  if server, e = ListenUDP("udp", a); e != nil {
    log.Fatalln("can't open socket for listening:", e.Error())
  }

  f(server)
}