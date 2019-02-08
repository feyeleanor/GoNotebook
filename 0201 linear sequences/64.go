package main
import . "fmt"
import "reflect"

type List struct { int; *List }

func main() {
  t := &List{4, &List{6, &List{8, nil}}}
  Range(struct { X, Y int; List }{0, 2, *t}, print_cell)
  Range(struct { *List; Y, X int }{ t, 2, 0 }, print_cell)
  Range(List{0, &List{2, t}}, print_cell)
}

func print_cell(i int, v interface{}) {
  switch s := v.(type) {
  case struct { X, Y int; List }:
    Printf("%v: (%v %v %v)\n", i, s.X, s.Y, s.List)
  case struct { *List; Y, X int }:
    print_cell(i, struct { X, Y int; List } { s.X, s.Y, *s.List })
  case List:
    Printf("%v: %v\n", i, s.int)
  case *List:
    print_cell(i, *s)
  default:
    Printf("%v: %v\n", i, v)
  }
}

func copy_value(v reflect.Value) reflect.Value {
  rv := reflect.New(v.Type()).Elem()
  rv.Set(v)
  return rv
}

func copy_data(t reflect.Type, s reflect.Value) (r reflect.Value) {
  r = copy_value(s)
  l := r.FieldByName("List")
  switch lt := l.Type(); {
  case !l.IsValid():
  case lt == reflect.PtrTo(t):
    l.Set(reflect.New(lt.Elem()))
  case lt == t:
    l.Set(reflect.New(lt).Elem())
  }
  return
}

func next_node(t reflect.Type, s reflect.Value) (r reflect.Value) {
  switch f := s.FieldByName("List"); f.Type() {
  case t, reflect.PtrTo(t):
    r = reflect.Indirect(f)
  }
  return
}

func range_list(s reflect.Value, f func(int, interface{})) {
  t := reflect.Indirect(s).Type()
  for i := 0; s.IsValid(); i++ {
    link := next_node(t, s)
    f(i, copy_data(t, s).Interface())
    s = link
  }
}

func Range(s interface{}, f func(int, interface{})) {
  range_list(reflect.ValueOf(s), f)
}
