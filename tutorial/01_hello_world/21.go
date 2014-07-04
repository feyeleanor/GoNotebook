package main

import "fmt"

type Hello struct {}

func (h Hello) String() string {
  return "Hello"
}

type Message struct {
  *Hello
  World string
}

func (v Message) String() (r string) {
  if v.Hello == nil {
    r = v.World
  } else {
    r = fmt.Sprintf("%v %v", v.Hello, v.World)
  }
  return
}

func main() {
  m := &Message{}
  fmt.Println(m)
  m.Hello = new(Hello)
  fmt.Println(m)
  m.World = "world"
  fmt.Println(m)
}