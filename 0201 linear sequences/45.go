package main
import . "fmt"
import "reflect"
import "unsafe"

var A [5]int
var B = []int{0, 2, 4, 6, 8}

func main() {
  print_array_from_slice("A", A)
  print_array_from_slice("B", B)
  A = array_from_slice(B)
  print_array_from_slice("A", A)
  A[2] = -1
  print_array_from_slice("A", A)
  print_array_from_slice("B", B)
}

func print_array_from_slice(label string, s interface{}) {
  Println(label, reflect.TypeOf(s), s)
}

func array_from_slice(s []int) [5]int {
  v := reflect.ValueOf(s)
  if !v.CanAddr() {
    x := reflect.New(v.Type()).Elem()
    x.Set(v)
    v = x
  }
  h := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
  return *(*[5]int)(unsafe.Pointer(h.Data))
}
