package main
import . "fmt"
import "reflect"

type Triple struct { L, M, R int }

func main() {
  t := Triple { 4, 6, 8 }
  Range(struct { X, Y int; List Triple }{ 0, 2, t }, print_cell)
  Range(struct { List *Triple; Y, X int }{ &t, 2, 0 }, print_cell)
}

func print_cell(i int, v interface{}) {
  switch s := v.(type) {
  case struct { X, Y int; List Triple }:
    Printf("%v: (%v %v)\n", i, s.X, s.Y)
  case struct { List *Triple; Y, X int }:
    print_cell(i, struct { X, Y int; List Triple } { s.X, s.Y, *s.List })
  default:
    Printf("%v: %v\n", i, v)
  }
}

func (t Triple) String() string {
  return Sprintf("(%v, %v, %v)", t.L, t.M, t.R)
}

func copy_value(v reflect.Value) reflect.Value {
  rv := reflect.New(v.Type()).Elem()
  rv.Set(v)
  return rv
}

func range_list(s reflect.Value, f func(int, reflect.Value)) {
  for i := 0; s.IsValid(); i++ {
    link := s.FieldByName("List")
    f(i, copy_value(s))
    s = reflect.Indirect(link)
  }
}

func Range(s interface{}, f func(int, interface{})) {
  range_list(reflect.ValueOf(s), func(i int, v reflect.Value) {
    f(i, v.Interface())
  })
}
