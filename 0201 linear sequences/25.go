package main
import . "fmt"
import "reflect"

type MyInt int
type SliceOfInt []int

func main() {
  print_with_status(SliceOfInt{0, -2, -4, -6, -8})
  print_with_status([]int{0, 2, 4, 6, 8})
  print_with_status([]uint{1, 3, 5, 7, 9})
  print_with_status([]MyInt{-1, -3, -5, -7, -9})
}

var T_SLICE = reflect.TypeOf([]int{})

func print_elements(s interface{}) bool {
  ok := true
  switch v := reflect.ValueOf(s); {
  case v.Type().ConvertibleTo(T_SLICE):
    for i, v := range v.Convert(T_SLICE).Interface().([]int) {
      Printf("%v: %v\n", i, v)
    }
  case v.Kind() == reflect.Slice:
    for i := 0; i < v.Len(); i++ {
      Printf("%v: %v\n", i, v.Index(i).Interface())
    }
  default:
    ok = false
  }
  return ok
}

func as_unnamed_slice(s []int) []int {
  return s
}

func print_with_status(s interface{}) {
  if b := print_elements(s); b {
    Println("succeeded")
  } else {
    Println("failed")
  }
}
