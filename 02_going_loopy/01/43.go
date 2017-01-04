package main
import . "fmt"

type Sequence interface {
	Next() (interface{}, Sequence)
}

type List struct { int; *List }
type SliceList []int

func main() {
	Range([]int{0, 2, 4, 6, 8}, print_cell)
	Range([]int{0, 2, 4, 6, 8}, func(i, v interface{}) {
		Printf("%v: %v\n", i, v)
	})
	Range(NewList(0, 2, 4, 6, 8), print_cell)
	Range(NewList(0, 2, 4, 6, 8), func(i, v interface{}) {
		Printf("%v: %v\n", i, v)
	})
	Range(SliceList{0, 2, 4, 6, 8}, print_cell)
	Range(SliceList{0, 2, 4, 6, 8}, func(i, v interface{}) {
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

func (l *List) Next() (v interface{}, r Sequence) {
	if l != nil {
		v = l.int
		r = l.List
	}
	return
}

func (s SliceList) Next() (v interface{}, r Sequence) {
	if len(s) > 0 {
		v = s[0]
		r = s[1:]
	}
	return
}

func Range(s, f interface{}) {
	switch s := s.(type) {
	case Sequence:
		if s != nil {
			i := 0
			switch f := f.(type) {
			case func(int, int):
				for v, r := s.Next(); r != nil; v, r = r.Next() {
					f(i, v.(int))
					i++
				}
			case func(interface{}, interface{}):
				for v, r := s.Next(); r != nil; v, r = r.Next() {
					f(i, v)
					i++
				}
			}
		}
	case []int:
		switch f := f.(type) {
		case func(int, int):
			for i, v := range s {
				f(i, v)
			}
		case func(interface{}, interface{}):
			for i, v := range s {
				f(i, v)
			}
		}
	}
}