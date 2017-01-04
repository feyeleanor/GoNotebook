package main
import (
  . "fmt"
  "os"
  "strings"
)

func main() {
  Println(strings.Join(os.Args[1:], " "))
}