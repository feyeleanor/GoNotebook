package main
import . "fmt"

type List struct { int; *List }

func main() {
  NewList(0, 2, 4, 6, 8).print()
}

func NewList(s ...int) (r *List) {
  for i := len(s) - 1; i > -1; i-- {
    r = &List{ s[i], r }
  }
  return
}

func (l *List) print() {
  for i := 0; l != nil; l = l.List {
    Printf("%v: %v\n", i, l.int)
    i++
  }
}
