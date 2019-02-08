package main
import . "fmt"
import "reflect"
import "unsafe"

var A = [...]int{0, 2, 4, 6, 8}
var B []int

func main() {
  Println(reflect.TypeOf(A), A)
  Println(reflect.TypeOf(B), B)
  B = slice_from_array(&A)
  Println(reflect.TypeOf(B), B)
}

var T_SLICE = reflect.TypeOf([]int{})

func slice_from_array(a *[5]int) []int {
  p := reflect.ValueOf(a)
  h := &reflect.SliceHeader{
    Data: p.Elem().UnsafeAddr(), Len: 5, Cap: 5,
  }
  s := reflect.NewAt(T_SLICE, unsafe.Pointer(h))
  return s.Elem().Interface().([]int)
}
