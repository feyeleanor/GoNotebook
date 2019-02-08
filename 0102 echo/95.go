package main
import "os"
import . "flag"
import . "fmt"
import "strconv"
import . "strings"

var name, spacer *string
var repeats int

func init() {
  def_n := os.Getenv("DEF_NAME")
  if len(def_n) == 0 {
    def_n = "world"
  }

leanpub-start-insert
  def_c, _ := strconv.Atoi(os.Getenv("DEF_REPS"))
leanpub-end-insert

  name = String("n", def_n, "n: name of person to greet")
  spacer = String("s", ",", "s: separator between name and message")
  IntVar(&repeats, "c", def_c, "c: number of times to display the message")
}

func main() {
  Parse()
  if message := Join(Args(), " "); len(message) > 0 {
    for ; repeats > 0; repeats-- {
      Printf("hello %v%v %v\n", *name, *spacer, message)
    }
  } else {
    for ; repeats > 0; repeats-- {
      Println("hello", *name)
    }
  }
}
