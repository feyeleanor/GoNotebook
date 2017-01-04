package main
import . "fmt"

type Enumerable interface {
	Range(func(interface{}, interface{}))
}

type List struct { int; *List }

func main() {
	f := func(i, v interface{}) {
		Printf("%v: %v\n", i, v)
	}
	Range([]int{0, 2, 4, 6, 8}, f)
	Range(NewList(0, 2, 4, 6, 8), f)
}

func NewList(s ...int) (r *List) {
	for i := len(s) - 1; i > -1; i-- {
		r = &List{ s[i], r }
	}
	return
}

func (l *List) Range(f func(interface{}, interface{})) {
	for i := 0; l != nil; l = l.List {
		f(i, l.int)
		i++
	}
}

func Range(s interface{}, f func(interface{}, interface{})) {
	switch s := s.(type) {
	case Enumerable:
		s.Range(f)
	case []int:
		for i, v := range s {
			f(i, v)
		}
	}
}