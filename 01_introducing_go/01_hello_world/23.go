package main

import "fmt"

type Hello struct {}

func (h Hello) String() string {
  return "Hello"
}

type World struct {}

func (w *World) String() string {
  return "world"
}

type Message struct {
  X fmt.Stringer
  Y fmt.Stringer
}

func (v Message) String() (r string) {
  switch {
  case v.X == nil && v.Y == nil:
  case v.X == nil:
    r = v.Y.String()
  case v.Y == nil:
    r = v.X.String()
  default:
    r = fmt.Sprintf("%v %v", v.X, v.Y)
  }
  return 
}

func main() {
  m := &Message{}
  fmt.Println(m)
  m.X = new(Hello)
  fmt.Println(m)
  m.Y = new(World)
  fmt.Println(m)
  m.Y = m.X
  fmt.Println(m)
  m = &Message{ X: new(World), Y: new(Hello) }
  fmt.Println(m)
  m.X, m.Y = m.Y, m.X
  fmt.Println(m)
}