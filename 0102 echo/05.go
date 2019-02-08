package main
import "flag"
import . "fmt"
import . "strings"

var suppress_newline bool

func init() {
  flag.BoolVar(&suppress_newline, "n", false, "n: suppress printing of newline")
}

func main() {
  flag.Parse()
  Print(Join(flag.Args(), " "))
  if !suppress_newline {
    Println()
  }
}
