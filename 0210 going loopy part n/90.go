package main
import . "fmt"
import . "reflect"

func main() {
  s := []int{0, 2, 4, 6, 8}
  print_values(s)
  print_values(func(i int) int {
    return s[i]
  })
}

func print_values(s interface{}) {
  switch s := ValueOf(s); s.Kind() {
  case Func:
    for i := 0; i < 5; i++ {
      p := []Value{ ValueOf(i) }
      Printf("%v: %v\n", i, s.Call(p)[0].Interface())
    }
  case Slice:
    for i := 0; i < s.Len(); i++ {
      Printf("%v: %v\n", i, s.Index(i).Interface())
    }
  }
}
