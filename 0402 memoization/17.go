package main

import "bufio"
import "fmt"
import "os"

const FILE LineFile = "17.txt"

type LineFile string

func (l LineFile) Append(s string) (e error) {
  var o *os.File

  o, e = os.OpenFile(string(l), os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
  if e == nil {
    defer o.Close()
    w := bufio.NewWriter(o)
    defer w.Flush()
    w.WriteString(s + "\n")
  }
  if e != nil {
    fmt.Println(l, "error writing", e.Error())
  }
  return
}

func main() {
    for i := 5; i > 0; i-- {
    FILE.Append(fmt.Sprint(i))
  }
}
