package main
import "fmt"

func main() {
  fmt.Println(message("world"))
}

func message(name string) (message string) {
  message = fmt.Sprintf("hello %v", name)
  return message
}