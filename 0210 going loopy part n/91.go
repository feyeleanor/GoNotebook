package main
import . "fmt"
import . "reflect"

func main() {
  s := []int{0, 2, 4, 6, 8}
  print_values(s)
  print_values(func(i int) int {
    return s[i]
  })
  print_values(func(i int) (v int, ok bool) {
    defer func() {
      ok = recover() == nil
    }()
    v = s[i]
    return
  })
}

func print_values(s interface{}) {
  switch s := ValueOf(s); s.Kind() {
  case Func:
    print_func(s)
  case Slice:
    for i := 0; i < s.Len(); i++ {
      print_value(i, s.Index(i))
    }
  }
}

func print_func(s Value) {
  switch s.Type().NumOut() {
  case 1:
    p := make([]Value, 1)
    for i := 0; i < 5; i++ {
      p[0] = ValueOf(i)
      print_value(i, s.Call(p)[0])
    }
  case 2:
    each_value(s, func(i int, v Value) {
      print_value(i, v)
    })
  default:
    panic(s.Interface())
  }
}

func each_value(s Value, f func(int, Value)) {
  var i int
  p := []Value{ ValueOf(0) }
  r := s.Call(p)
  for r[1].Bool() {
    f(i, r[0])
    i++
    p[0] = ValueOf(i)
    r = s.Call(p)
  }
}

func print_value(i int, v Value) {
  Printf("%v: %v\n", i, v.Interface())
}
