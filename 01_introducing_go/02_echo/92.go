package main
import (
  "os"
  . "flag"
  . "fmt"
  . "strings"
)

var name, spacer *string
var repeats int

func init() {
leanpub-start-insert
  var def_n string

  if def_n = os.Getenv("DEF_NAME"); len(def_n) == 0 {
    def_n = "world"
  }

  name = String("n", def_n, "n: name of person to greet")
leanpub-end-insert
  spacer = String("s", ",", "s: separator between name and message")
  IntVar(&repeats, "c", 0, "c: number of times to display the message")
}

func main() {
  Parse()
  if message = Join(Args(), " "); len(message) > 0 {
    for ; repeats > 0; repeats-- {
      Printf("hello %v%v %v\n", *name, *spacer, message)
    }
  } else {
    for ; repeats > 0; repeats-- {
      Println("hello", *name)
    }
  }
}