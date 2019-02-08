package main
import . "fmt"

const Hello = "hello"
var world  string

func init() {
  world = "world"
  Println(Hello, world)
}

func main() {}