package main
import . "fmt"

const Hello = "hello"
var world  string

func init() {
  Print(Hello, " ")
  world = "world"
}

func init() {
  Printf("%v\n", world)
}

func main() {}