package main
import . "fmt"

type IntSlice []int
type SliceOfInt []int

func main() {
	s1 := IntSlice{9, 7, 5, 3, 1}
	s2 := SliceOfInt{0, -2, -4, -6, -8}
	s3 := []int{0, 2, 4, 6, 8}

	print_elements(s1)
	print_elements(s2)
	print_elements(s3)
}

func print_elements(s interface{}) {
	switch s := s.(type) {
	case []int:
		for i, v := range s {
			Printf("%v: %v\n", i, v)
		}
	case interface { print_elements () }:
		s.print_elements()
	}
}

func (i IntSlice) print_elements() {
	print_elements(as_unnamed_slice(i))
}

func (i SliceOfInt) print_elements() {
	print_elements(as_unnamed_slice(i))
}

func as_unnamed_slice(s []int) []int {
	return s
}