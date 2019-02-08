package main
import . "fmt"
import "reflect"

type IntSlice []int
type SliceOfInt []int

func main() {
  print_with_status(IntSlice{9, 7, 5, 3, 1})
  print_with_status(SliceOfInt{0, -2, -4, -6, -8})
  print_with_status([]int{0, 2, 4, 6, 8})
  print_with_status([]uint{1, 3, 5, 7, 9})
}

type ElementPrinter interface {
  print_elements() bool
}

func (i IntSlice) print_elements() {
  print_elements(as_unnamed_slice(i))
}

var T_SLICE = reflect.TypeOf([]int{})

func print_elements(s interface{}) bool {
  ok := true
  switch s := s.(type) {
  case []int:
    for i, v := range s {
      Printf("%v: %v\n", i, v)
    }
  case ElementPrinter:
    s.print_elements()
  default:
    v := reflect.ValueOf(s)
    if ok = v.Type().ConvertibleTo(T_SLICE); ok {
      print_elements((v.Convert(T_SLICE).Interface().([]int)))
    }
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
