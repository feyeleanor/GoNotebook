package main
import . "fmt"
import "reflect"

func main(){
  print_array([4]int{0, 2, 4, 6})
  print_array([5]int{0, 2, 4, 6, 8})
  print_array([6]int{0, 2, 4, 6, 8, 10}
)

func print_array(s interface{}) {
  if v := reflect.ValueOf(s); v.Kind() == reflect.Array {
    for i := 0; i < v.Len(); i++ {
      Printf("%v: %v\n", i, v.Index(i))
    }
  }
}
