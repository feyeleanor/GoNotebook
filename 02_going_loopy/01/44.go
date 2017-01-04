package main
import . "fmt"
import "reflect"

type List struct { Value int; *List }

func main() {
	Range(NewList(0, 2, 4, 6, 8), print_cell)
	Range(NewList(0, 2, 4, 6, 8), func(i, v interface{}) {
		Printf("%v: %v\n", i, v)
	})
}

func print_cell(i, v int) {
	Printf("%v: %v\n", i, v)
}

func NewList(s ...int) (r *List) {
	for i := len(s) - 1; i > -1; i-- {
		r = &List{ s[i], r }
	}
	return
}

func range_list(s reflect.Value, f func(int, reflect.Value)) {
	for i := 0; s.IsValid(); i++ {
		f(i, s.FieldByName("Value"))
		s = reflect.Indirect(s.FieldByName("List"))
	}
}

func Range(s, f interface{}) {
	switch s := reflect.Indirect(reflect.ValueOf(s)); s.Kind() {
	case reflect.Ptr:
		Range(s, f)
	case reflect.Struct:
		switch f := f.(type) {
		case func(int, int):
			range_list(s, func(i int, v reflect.Value) {
				f(i, int(v.Int()))
			})
		case func(interface{}, interface{}):
			range_list(s, func(i int, v reflect.Value) {
				f(i, v.Interface())
			})
		}
	}
}