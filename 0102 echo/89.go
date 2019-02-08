package main
import "flag"
import . "fmt"
import . "strings"

var name, spacer *string
var repeats int

func init() {
  name = flag.String("n", "world", "n: name of person to greet")
  spacer = flag.String("s", ",", "s: separator between name and message")
  flag.IntVar(&repeats, "c", 1, "c: number of times to display the message")
}

func main() {
  flag.Parse()
  message := Join(flag.Args(), " ")
  for i := repeats; i > 0; i-- {
    Printf("hello %v%v %v\n", *name, *spacer, message)
  }
}
