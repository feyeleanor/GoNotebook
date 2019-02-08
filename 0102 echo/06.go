package main
import "flag"
import . "fmt"
import . "strings"

var suppress_newline bool
leanpub-start-insert
var repetitions int
leanpub-end-insert

func init() {
  flag.BoolVar(&suppress_newline, "n", false, "n: suppress printing of newline")
leanpub-start-insert
  flag.IntVar(&repetitions, "r", 1, "r: repeat the message more than once")
leanpub-end-insert
}

func main() {
  flag.Parse()
  message := Join(flag.Args(), " ")
leanpub-start-insert
  for i := 0; i < repetitions; i++ {
    Print(message)
  }
leanpub-end-insert
  if !suppress_newline {
    Println()
  }
}
