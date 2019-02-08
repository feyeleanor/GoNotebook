package main
import . "fmt"

type List struct { int; *List }

func main() {
  NewList(0, 2, 4, 6, 8).Range(func(i, v int) {
    Printf("%v: %v\n", i, v)
  })
}

func NewList(s ...int) (r *List) {
  for i := len(s) - 1; i > -1; i-- {
    r = &List{ s[i], r }
  }
  return
}

func (l *List) Range(f func(int, int)) {
  for i := 0; l != nil; l = l.List {
    f(i, l.int)
    i++
  }
}
