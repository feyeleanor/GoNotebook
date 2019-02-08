package main
import . "fmt"
import "reflect"

func main() {
  print_array([...]interface{}{0, 1.0, 'a', "a", 4})
}

func print_array(s [5]interface{}) {
  v := reflect.ValueOf(s)
  for i := 0; i < 5; i++ {
    if x, ok := v.Index(i).Interface().(int); ok {
      Printf("%v: %v\n", i, x)
    }
  }
}
