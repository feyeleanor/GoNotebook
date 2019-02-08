package main
import "flag"
import . "fmt"
import . "strings"

var suppress_newline *bool

leanpub-start-insert
func init() {
  suppress_newline = flag.Bool("n", false, "n: suppress printing of newline")
}
leanpub-end-insert

func main() {
leanpub-start-insert
  flag.Parse()
  Print(Join(flag.Args(), " "))
leanpub-end-insert
  if ! *suppress_newline {
    Println()
  }
}
