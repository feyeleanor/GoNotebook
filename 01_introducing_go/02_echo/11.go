package main
import (
  "flag"
  . "fmt"
  "os"
  . "strings"
)

var suppress_newline bool
var repetitions int

func init() {
  flag.BoolVar(&suppress_newline, "n", false, "n: suppress printing of newline")
  flag.IntVar(&repetitions, "r", 1, "r: repeat the message more than once")
}

func main() {
  flag.Parse()
  if repetitions < 0 {
    os.Exit(1)
  }
  message := Join(flag.Args(), " ")
  if suppress_newline {
    repeat(func () {
      Print(message)
    })
  } else {
    repeat(func () {
      Println(message)
    })
  }
}

func repeat(f func()) {
  for ; repetitions > 0; repetitions-- {
    f()
  }
}