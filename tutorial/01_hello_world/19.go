package main

import "fmt"

type HelloWorld struct {}

func (h HelloWorld) String() string {
  return "Hello world"
}

type Message struct {
  HelloWorld
}

func main() {
  m := &Message{}
  fmt.Println(m.String())
  fmt.Println(m)
}