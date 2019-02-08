package main
import . "fmt"
import "reflect"

func main() {
  print_array([4]int8{0, 2, 4, 6})
  print_array([5]int{0, 2, 4, 6, 8})
  print_array([6]uint{0, 2, 4, 6, 8, 10})
  print_array([4]interface{}{0, 1.0, 'a', "a"})
}

var T_INT = reflect.TypeOf(int(0))

func print_array(s interface{}) {
  v := reflect.ValueOf(s)
  t := v.Type().Elem()
  if v.Kind() == reflect.Array && t.AssignableTo(T_INT) {
    for i := 0; i < v.Len(); i++ {
      Printf("%v: %v\n", i, v.Index(i))
    }
  }
}
