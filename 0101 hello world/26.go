package main
import . "fmt"

const Hello = "hello"
var world  string

func init() {
  world = "world"
}

func main() {
  Println(Hello, world)
}