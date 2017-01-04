package main

import (
  "bytes"
  "crypto/rand"
  "crypto/rsa"
  "crypto/sha1"
  "crypto/x509"
  "encoding/gob"
  "encoding/pem"
  "io/ioutil"
  . "fmt"
  . "net"
)

var RSA_LABEL = []byte("served")

func main() {
  Connect(":1025", func(server *UDPConn, private_key *rsa.PrivateKey) {
    cipher_text := MakeBuffer()
    if n, e := server.Read(cipher_text); e == nil {
      if plain_text, e := rsa.DecryptOAEP(sha1.New(), rand.Reader, private_key, cipher_text[:n], RSA_LABEL); e == nil {
        Println((string)(plain_text))
      }
    }
  })
}

func Connect(address string, f func(*UDPConn, *rsa.PrivateKey)) {
  LoadPrivateKey("client.key.pem", func(private_key *rsa.PrivateKey) {
    if address, e := ResolveUDPAddr("udp", ":1025"); e == nil {
      if server, e := DialUDP("udp", nil, address); e == nil {
        defer server.Close()
        SendKey(server, private_key.PublicKey, func() {
          f(server, private_key)
        })
      }
    }
  })
}

func LoadPrivateKey(file string, f func(*rsa.PrivateKey)) {
  if file, e := ioutil.ReadFile(file); e == nil {
    if block, _ := pem.Decode(file); block != nil {
      if block.Type == "RSA PRIVATE KEY" {
        if key, _ := x509.ParsePKCS1PrivateKey(block.Bytes); key != nil {
          f(key)
        }
      }
    }
  }
  return
}

func SendKey(server *UDPConn, public_key rsa.PublicKey, f func()) {
  var encoded_key bytes.Buffer
  if e := gob.NewEncoder(&encoded_key).Encode(public_key); e == nil {
    if _, e = server.Write(encoded_key.Bytes()); e == nil {
      f()
    }
  }
}

func MakeBuffer() (r []byte) {
  return make([]byte, 1024)
}