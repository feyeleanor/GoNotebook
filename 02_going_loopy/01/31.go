package main
import . "fmt"

type SparseArray map[int] int

func main() {
	s := SparseArray{2: 4, 4: 8, 6: 12, 8: 16}
	print_sparse_array(s)
}

func print_sparse_array(s SparseArray) {
	s.Range(func(i, v int) {
		Printf("%v: %v\n", i, v)
	})
}

func (s SparseArray) Range(f func(int, int)) {
	var v int
	var ok bool
	for i, n := 0, len(s); n > 0; i++ {
		if v, ok = s[i]; ok {
			n--
		}
		f(i, v)
	}
}