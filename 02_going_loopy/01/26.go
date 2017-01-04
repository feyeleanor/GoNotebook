package main
import . "fmt"

type IntSlice []int
type SliceOfInt []int
type IntMap map[int] int
type MapOfInt map[int] int

func main() {
	s1 := IntSlice{9, 7, 5, 3, 1}
	s2 := SliceOfInt{0, -2, -4, -6, -8}
	s3 := []int{0, 2, 4, 6, 8}

	m1 := IntMap{0: 9, 1: 7, 2: 5, 3: 3, 4: 1}
	m2 := MapOfInt{0: 0, 1: -2, 2: -4, 3: -6, 4: -8}
	m3 := map[int] int{0: 0, 1: 2, 2: 4, 3: 6, 4: 8}

	for _, v := range []interface{} {s1, s2, s3, m1, m2, m3} {
		print_elements(v)
	}
}

func print_elements(V interface{}) {
	switch V := V.(type) {
	case []int:
		for i, v := range V {
			Printf("%v: %v\n", i, v)
		}
	case IntSlice:
		print_elements(as_unnamed_slice(V))
	case SliceOfInt:
		print_elements(as_unnamed_slice(V))
	case map[int] int:
		for i, v := range V {
			Printf("%v: %v\n", i, v)
		}
	case IntMap:
		print_elements(as_unnamed_map(V))
	case MapOfInt:
		print_elements(as_unnamed_map(V))
	}
}

func as_unnamed_slice(s []int) []int {
	return s
}

func as_unnamed_map(m map[int] int) map[int] int {
	return m
}