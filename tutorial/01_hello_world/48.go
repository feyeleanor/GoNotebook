package main

import (
  "bytes"
  "crypto/rand"
  "crypto/rsa"
  "crypto/sha1"
  "encoding/gob"
  . "fmt"
  . "net"
  "os"
)

var HELLO_WORLD = []byte("Hello World")
var RSA_LABEL = []byte("served")

func main() {
  Serve(":1025", func(s *UDPConn, c *UDPAddr, packet *bytes.Buffer) {
    var e error
    var n int
    var key rsa.PublicKey
    var response []byte

    if e = gob.NewDecoder(packet).Decode(&key); e != nil {
      Abort("unable to decode key:", packet)
    }

    if response, e = rsa.EncryptOAEP(sha1.New(), rand.Reader, &key, HELLO_WORLD, RSA_LABEL); e != nil {
      Abort("unable to encrypt key")
    }

    if n, e = s.WriteToUDP(response, c); e != nil {
      Abort("unable to write response")
    }
    Printf("server %v bytes written to %v\n", n, c)
  })
}

func Serve(address string, f func(*UDPConn, *UDPAddr, *bytes.Buffer)) {
  var e error
  var n int
  var a, client *UDPAddr
  var server *UDPConn

  if a, e = ResolveUDPAddr("udp", address); e != nil {
    Abort("unable to resolve UDP address")
  }

  if server, e = ListenUDP("udp", a); e != nil {
    Abort("can't open socket for listening")
  }

  for buffer := MakeBuffer(); ; buffer = MakeBuffer() {
    if n, client, e = server.ReadFromUDP(buffer); e != nil {
      Abort("unable to read from client")
    }
    go f(server, client, bytes.NewBuffer(buffer[:n]))
  }
}

func MakeBuffer() (r []byte) {
  return make([]byte, 1024)
}

func Abort(data ...interface{}) {
  Println(data...)
  os.Exit(1)
}