package main
import . "fmt"
import "reflect"

type IntMap map[int] int
type MyInt int
type SparseArray struct {
	m map[int] int
	Count int
}

func main() {
	print_elements([]MyInt{-9, -7, -5, -3, -1})
	print_elements(map[int] int{0: 0, 1: 2, 2: 4, 3: 6, 4: 8})
	print_elements(map[int] MyInt{0: 0, 1: 2, 2: 4, 3: 6, 4: 8})
	print_elements(IntMap{0: 0, 1: 2, 2: 4, 3: 6, 4: 8})
	print_elements(NewSparseArray(map[int] int{2: 4, 4: 8}))
}

type ElementPrinter interface {
	print_elements ()
}

func NewSparseArray(m map[int] int) (r *SparseArray) {
	r = &SparseArray{m: make(map[int] int)}
	for k, v := range m {
		r.m[k] = v
		if k > r.Count {
			r.Count = k
		}
	}
	return
}

func (s SparseArray) Range(f func(int, int)) {
	for i := 0; s.Count > -1; i++ {
		f(i, s.m[i])
		s.Count--
	}
}

func print_elements(m interface{}) {
	switch m := m.(type) {
	case SparseArray:
		m.Range(func(k, v int) {
			Printf("%v: %v\n", k, v)
		})
	case *SparseArray:
		print_elements(*m)
	default:
		switch v := reflect.ValueOf(m); v.Kind() {
		case reflect.Slice:
			defer func() {
				recover()
			}()
			for i := 0; ; i++ {
				Printf("%v: %v\n", i, v.Index(i).Interface())
			}
		case reflect.Map:
			for i, k := range v.MapKeys() {
				Printf("%v: %v\n", i, v.MapIndex(k).Interface())
			}
		default:
			panic(m)
		}
	}
}