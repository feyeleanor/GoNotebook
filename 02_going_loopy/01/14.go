package main
import . "fmt"

type IntSlice []int
type SliceOfInt []int

func main() {
	print_slice([]int{0, 2, 4, 6, 8})
	print_slice(IntSlice{9, 7, 5, 3, 1})
	print_slice(SliceOfInt{0, -2, -4, -6, -8})

	print_IntSlice([]int{0, 2, 4, 6, 8})
	print_IntSlice(IntSlice{9, 7, 5, 3, 1})

	print_SliceOfInt([]int{0, 2, 4, 6, 8})
	print_SliceOfInt(SliceOfInt{0, -2, -4, -6, -8})
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