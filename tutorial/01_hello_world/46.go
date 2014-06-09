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
  Serve(":1025", func(s *UDPConn, c *UDPAddr, packet *bytes.Buffer) {
    var key rsa.PublicKey
    if e := gob.NewDecoder(packet).Decode(&key); e == nil {
      if response, e := rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e == nil {
        if n, e := s.WriteToUDP(response, c); e == nil {
          Printf("server %v bytes written to %v\n", n, c)
        }
      }
    }
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer)) {
  if address, e := ResolveUDPAddr("udp", address); e == nil {
    if server, e := ListenUDP("udp", address); e == nil {
      for buffer := MakeBuffer(); ; buffer = MakeBuffer() {
        if n, client, e := server.ReadFromUDP(buffer); e == nil {
          go f(server, client, bytes.NewBuffer(buffer[:n]))
        }
      }
    }
  }
}

func MakeBuffer() (r []byte) {
  return make([]byte, 1024)
}