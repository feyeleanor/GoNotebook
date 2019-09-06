package main

import "bufio"
import "fmt"
import "os"

const FILE LineFile = "18.txt"

type LineFile string

func (l LineFile) Append(f func(*bufio.Writer) error) (e error) {
  var o *os.File

  o, e = os.OpenFile(string(l), os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
  if e == nil {
    defer o.Close()
    w := bufio.NewWriter(o)
    defer w.Flush()
    e = f(w)
  }
  if e != nil {
    fmt.Println(l, "error writing", e.Error())
  }
  return
}

func main() {
  FILE.Append(func(w *bufio.Writer) (e error) {
    for i := 5; i > 0 && e == nil; i-- {
      _, e = w.WriteString(fmt.Sprint(i))
    }
    return
  })
}
