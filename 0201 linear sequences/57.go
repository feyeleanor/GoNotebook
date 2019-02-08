package main
import . "fmt"
import "reflect"

type List struct { Value int; *List }

func main() {
  Range(NewList(0, 2, 4, 6, 8), print_cell)
  Range(NewList(0, 2, 4, 6, 8), func(i, v interface{}) {
    Printf("%v: %v\n", i, v)
  })
  Range(struct { List; X int }{ *NewList(2, 4, 6, 8), 0 }, print_cell)
  Range(struct { X int; *List }{ 0, NewList(2, 4, 6, 8) }, print_cell)
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

func range_list(s reflect.Value, f func(int, reflect.Value)) {
  for i := 0; s.IsValid(); i++ {
    if t := s.Type(); t.Field(0).Name == "List" {
      f(i, s.Field(1))
      s = s.Field(0)
    } else {
      f(i, s.Field(0))
      s = s.Field(1)
    }
  s = reflect.Indirect(s)
  }
}

func Range(s, f interface{}) {
  switch s := reflect.Indirect(reflect.ValueOf(s)); s.Kind() {
  case reflect.Ptr:
    Range(s, f)
  case reflect.Struct:
    switch f := f.(type) {
    case func(int, int):
      range_list(s, func(i int, v reflect.Value) {
        f(i, int(v.Int()))
      })
    case func(interface{}, interface{}):
      range_list(s, func(i int, v reflect.Value) {
        f(i, v.Interface())
      })
    }
  }
}
