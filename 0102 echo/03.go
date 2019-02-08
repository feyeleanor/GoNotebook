package main
import . "fmt"
import "os"
import "strings"

func main() {
  Println(strings.Join(os.Args[1:], " "))
}
