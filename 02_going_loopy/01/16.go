package main
import . "fmt"

type IntSlice []int
type SliceOfInt []int

func main() {
	s1 := IntSlice{9, 7, 5, 3, 1}
	s2 := SliceOfInt{0, -2, -4, -6, -8}
	s3 := []int{0, 2, 4, 6, 8}

	print_slice(s1)
	print_slice(s2)
	print_slice(s3)

	print_IntSlice(s1)
	print_IntSlice(as_unnamed_slice(s2))
	print_IntSlice(s3)

	print_SliceOfInt(as_unnamed_slice(s1))
	print_SliceOfInt(s2)
	print_SliceOfInt(s3)
}

func print_slice(s []int) {
	for i, v := range s {
		Printf("%v: %v\n", i, v)
	}
}

func print_IntSlice(s IntSlice) {
	for i, v := range s {
		Printf("%v: %v\n", i, v)
	}
}

func print_SliceOfInt(s SliceOfInt) {
	for i, v := range s {
		Printf("%v: %v\n", i, v)
	}
}

func as_unnamed_slice(s []int) []int {
	return s
}