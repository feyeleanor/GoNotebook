package main
import . "fmt"
import "reflect"

func main() {
  print_array([]int{0, 2, 4, 6})
  print_array([5]int{0, 2, 4, 6, 8})
  print_array([5]interface{}{0, 1.0, 'a', "a", 4})
  print_array([]interface{}{0, 1.0, 'a', "a", 4})
}

var T_INT = reflect.TypeOf(int(0))

func print_array(s interface{}) {
  defer func() {
    recover()
  }()
  v := reflect.ValueOf(s)
  switch v.Kind() {
  case reflect.Array, reflect.Slice:
    for i := 0; i < v.Len(); i++ {
      if x, ok := v.Index(i).Interface().(int); ok {
        Printf("%v: %v\n", i, x)
      }
    }
  }
}
