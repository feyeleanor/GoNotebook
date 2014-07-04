package main

import (
  "bytes"
  "crypto/rand"
  "crypto/rsa"
  "crypto/sha1"
  "encoding/gob"
  . "fmt"
  . "net"
)

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

func main() {
  Serve(":1025", func(connection *UDPConn, c *UDPAddr, packet *bytes.Buffer) {
    var key rsa.PublicKey
    if e := gob.NewDecoder(packet).Decode(&key); e == nil {
      if response, e := rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e == nil {
        if n, e := connection.WriteToUDP(response, c); e == nil {
          Println(n, "bytes written to", c)
        }
      }
    }
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer)) {
  Launch(address, func(connection *UDPConn) {
    for {
      buffer := make([]byte, 1024)
      if n, client, e := connection.ReadFromUDP(buffer); e == nil {
        go f(connection, client, bytes.NewBuffer(buffer[:n]))
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