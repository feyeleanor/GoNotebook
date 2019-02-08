package main
import . "fmt"

type Iteration struct {
  i int
  v interface{}
}

type Enumerable interface {
  Enumerator() chan Iteration
}

type List struct {
  int
  *List
}

func main() {
  for x := range Enumerator([]int{0, 2, 4, 6, 8}) {
    Printf("%v: %v\n", x.i, x.v)
  }
  for x := range Enumerator(NewList(0, 2, 4, 6, 8)) {
    Printf("%v: %v\n", x.i, x.v)
  }
}

func NewList(s ...int) (r *List) {
  for i := len(s) - 1; i > -1; i-- {
    r = &List{ s[i], r }
  }
  return
}

func (l *List) Enumerator() (r chan Iteration) {
  r = make(chan Iteration)
  go func() {
    for i := 0; l != nil; i++ {
      r <- Iteration{ i, l.int }
      l = l.List
    }
    close(r)
  }()
  return
}

func Enumerator(s interface{}) (r chan Iteration) {
  switch s := s.(type) {
  case Enumerable:
    r = s.Enumerator()
  case []int:
    r = make(chan Iteration)
    go func() {
      for i, v := range s {
        r <- Iteration{ i, v }
      }
      close(r)
    }()
  }
  return
}
