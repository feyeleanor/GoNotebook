package main
import . "fmt"
import "reflect"

type IntSlice []int
type SliceOfInt []int

func main() {
  print_elements(IntSlice{9, 7, 5, 3, 1})
  print_elements(SliceOfInt{0, -2, -4, -6, -8})
  print_elements([]int{0, 2, 4, 6, 8})
}

type ElementPrinter interface {
  print_elements ()
}

func (i IntSlice) print_elements() {
  print_elements(as_unnamed_slice(i))
}

var T_SLICE reflect.Type

func init() {
  T_SLICE = reflect.TypeOf([]int{})
}

func print_elements(s interface{}) {
  switch s := s.(type) {
  case []int:
    for i, v := range s {
      Printf("%v: %v\n", i, v)
    }
  case ElementPrinter:
    s.print_elements()
  default:
    v := reflect.ValueOf(s)
    if v.Type().ConvertibleTo(T_SLICE) {
      print_elements((v.Convert(T_SLICE).Interface().([]int)))
    } else {
      panic(s)
    }
  }
}

func as_unnamed_slice(s []int) []int {
  return s
}
