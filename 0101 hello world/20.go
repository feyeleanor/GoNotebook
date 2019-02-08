package main
import "fmt"

type HelloWorld bool

func (h HelloWorld) String() (r string) {
  if h {
    r = "Hello world"
  }
  return
}

type Message struct {
  HelloWorld
}

func main() {
  m := &Message{ HelloWorld: true }
  fmt.Println(m)
  m.HelloWorld = false
  fmt.Println(m)
  m.HelloWorld = true
  fmt.Println(m)
}
