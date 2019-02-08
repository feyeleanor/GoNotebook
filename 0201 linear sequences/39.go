package main
import . "fmt"
import "reflect"

func main() {
  print_array(37)
  print_array([]int{0, 2, 4, 6})
  print_array([5]int{0, 2, 4, 6, 8})
  print_array([6]uint{0, 2, 4, 6, 8, 10})
  print_array([]interface{}{0, 1.0, 'a', "a"})
}

var T_INT = reflect.TypeOf(int(0))

func print_array(s interface{}) {
  defer func() {
    recover()
  }()
  v := reflect.ValueOf(s)
  if t := v.Type().Elem(); t.AssignableTo(T_INT) {
    switch v.Kind() {
    case reflect.Array, reflect.Slice:
      for i := 0; i < v.Len(); i++ {
        Printf("%v: %v\n", i, v.Index(i))
      }
    }
  }
}
