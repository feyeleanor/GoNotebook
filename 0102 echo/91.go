package main
import "flag"
import . "fmt"
import . "strings"
import "os"

var name, spacer *string
var repeats int

func init() {
  name = flag.String("n", "world", "n: name of person to greet")
  spacer = flag.String("s", ",", "s: separator between name and message")
  flag.IntVar(&repeats, "c", 1, "c: number of times to display the message")
}

func main() {
leanpub-start-insert
  flag.Parse()
  if repeats < 0 {
    os.Exit(1)
  }
  if message = Join(flag.Args(), " "); len(message) > 0 {
leanpub-end-insert
    for ; repeats > 0; repeats-- {
      Printf("hello %v%v %v\n", *name, *spacer, message)
    }
  } else {
    for ; repeats > 0; repeats-- {
      Println("hello", *name)
    }
  }
}
