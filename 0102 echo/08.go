package main
import "flag"
import . "fmt"
import "os"
import . "strings"

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
  for i := 0; i < repetitions; i++ {
    if suppress_newline {
      Print(message)
    } else {
      Println(message)
    }
  }
}
