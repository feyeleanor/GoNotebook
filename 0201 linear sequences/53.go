package main
import . "fmt"

type Enumerable interface {
  Range(interface{})
}

type List struct { int; *List }

func main() {
  Range([]int{0, 2, 4, 6, 8}, print_cell)
  Range([]int{0, 2, 4, 6, 8}, func(i, v interface{}) {
    Printf("%v: %v\n", i, v)
  })
  Range(NewList(0, 2, 4, 6, 8), print_cell)
  s := make([]int, 0, 5)
  Range(NewList(0, 2, 4, 6, 8), &s)
  Range(s, print_cell)
}

func print_cell(i, v int) {
  Printf("%v: %v\n", i, v)
}

func NewList(s ...int) (r *List) {
  for i := len(s) - 1; i > -1; i-- {
    r = &List{ s[i], r }
  }
  return
}

func (l *List) Range(f interface{}) {
  switch f := f.(type) {
  case func(int, int):
    for i := 0; l != nil; l = l.List {
      f(i, l.int)
      i++
    }
  case *[]int:
    for i := 0; l != nil && i < cap(*f); l = l.List {
      if len(*f) - 1 < i {
        *f = append(*f, l.int)
      } else {
        (*f)[i] = l.int
      }
      i++
    }
  }
}

func Range(s, f interface{}) {
  switch s := s.(type) {
  case Enumerable:
    s.Range(f)
  case []int:
    switch f := f.(type) {
    case func(int, int):
      for i, v := range s {
        f(i, v)
      }
    case func(interface{}, interface{}):
      for i, v := range s {
        f(i, v)
      }
    }
  }
}
