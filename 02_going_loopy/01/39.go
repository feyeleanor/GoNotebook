package main
import . "fmt"

type Enumerable interface {
	Range(func(int, int))
}

type IntSlice []int
type MyIntSlice struct { IntSlice }
type List struct { int; *List }


func main() {
	Range([]int{0, 2, 4, 6, 8}, print_cell)
	Range(IntSlice{0, 2, 4, 6, 8}, print_cell)
	Range(([]int)(IntSlice{0, 2, 4, 6, 8}), print_cell)
	Range(MyIntSlice{IntSlice{0, 2, 4, 6, 8}}, print_cell)
	Range(NewList(0, 2, 4, 6, 8), print_cell)
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

func (l *List) Range(f func(int, int)) {
	for i := 0; l != nil; l = l.List {
		f(i, l.int)
		i++
	}
}

func (m MyIntSlice) Range(f func(int, int)) {
	for i, v := range m.IntSlice {
		f(i, v)
	}
}

func Range(s interface{}, f func(int, int)) {
	switch s := s.(type) {
	case Enumerable:
		s.Range(f)
	case []int:
		for i, v := range s {
			f(i, v)
		}
	}
}