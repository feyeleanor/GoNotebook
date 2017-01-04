package main
import . "fmt"
import "reflect"

type Pair struct { L, R int }

func main() {
	Range(struct { X int; List Pair }{ 0, Pair{ 2, 4 } }, print_cell)
	Range(struct { List *Pair; X int }{ &Pair{ 2, 4 }, 0 }, print_cell)
}

func print_cell(i int, v interface{}) {
	Printf("%v: %v\n", i, v)
}

func range_list(s reflect.Value, f func(int, reflect.Value)) {
	for i := 0; s.IsValid(); i++ {
		switch t := s.Type(); {
		case t.Field(0).Name == "List":
			f(i, s.Field(1))
			s = s.Field(0)
		case t.Field(1).Name == "List":
			f(i, s.Field(0))
			s = s.Field(1)
		default:
			f(i, s)
			s = reflect.ValueOf(nil)
		}
		s = reflect.Indirect(s)
	}
}

func Range(s interface{}, f func(int, interface{})) {
	range_list(reflect.ValueOf(s), func(i int, v reflect.Value) {
		f(i, v.Interface())
	})
}