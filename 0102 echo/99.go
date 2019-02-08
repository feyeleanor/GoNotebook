package main
import "os"
import . "flag"
import "fmt"
import "io/ioutil"
import "strconv"
import . "strings"

var name, spacer *string
var repeats int

func init() {
leanpub-start-insert
  name = String("n", defaultName(), "n: name of person to greet")
leanpub-end-insert
  spacer = String("s", ",", "s: separator between name and message")
  file := String("f", "", "f: name of a file containing a block of text to display")
leanpub-start-insert
  IntVar(&repeats, "c", defaultRepeats(), "c: number of times to display the message")
leanpub-end-insert
}

func main() {
  message, text_block := readFlags()
  switch {
  case len(message) > 0:
    message = fmt.Sprintf("hello %v%v %v", *name, *spacer, message)
    if len(text_block) > 0 {
      message += "\n" + text_block
    }
  case len(text_block) > 0:
    message = fmt.Sprintf("hello %v\n%v", *name, text_block)
  default:
    message = fmt.Sprintf("hello %v", *name)
  }

  for ; repeats > 0; repeats-- {
    fmt.Println(message)
  }
}

func readFlags() (message, text_block string) {
  Parse()
  return loadMessage(*file), Join(Args(), " ")
}

leanpub-start-insert
func defaultName() (r string) {
  if r = os.Getenv("DEF_NAME"); len(r) == 0 {
    r = "world"
  }
  return
}

func defaultRepeats() (r int) {
  if d := os.Getenv("DEF_REPS"); len(d) != 0 {
    var e error
    if r, e = strconv.Atoi(d); e != nil {
      fmt.Println("DEF_REPS must be an integer:", d)
    }
  }
  return
}

func loadMessage(filename string) (r string) {
  if len(filename) > 0 {
    if text, err := ioutil.ReadFile(filename); err == nil {
      r = string(text)
    } else {
      fmt.Println("no such file:", filename)
    }
  }
  return
}
leanpub-end-insert
