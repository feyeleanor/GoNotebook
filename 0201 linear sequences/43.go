package main
import . "fmt"
import "reflect"

var A int
var B = [...]int{0, 2, 4, 6, 8}
var C = []int{1, 3, 5, 7}

func main() {
  var D map[int] int

  check_addressability("LITERAL VALUES", [...]int{0, 1}, []int{1, 3})
  check_addressability("INTEGER VARIABLES", A, &A)
  check_addressability("ARRAY VARIABLES", B, B[0], &B, &B[0])
  check_addressability("SLICE VARIABLES", C, C[0], &C, &C[0])
  check_addressability("MAP VARIABLES", D, D[0], &D)
}

func check_addressability(s string, v ...interface{}) {
  Printf("\t%v\n", s)
  for _, v := range v {
    check_addressability_with_reflection(v)
  }
  Println()
}

func check_addressability_with_reflection(s interface{}) {
  v := reflect.ValueOf(s)
  t := Sprint("\t", v.Type(), v)
  Println(t, "is addressable?", v.CanAddr())

  defer func() {
    recover()
  }()
  x := v.Elem()
  Println(t, "content", x.Type(), "is addressable?", x.CanAddr())
}
