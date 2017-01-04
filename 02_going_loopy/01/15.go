package main

type IntSlice []int
type SliceOfInt []int

func main() {
	var s1 []int
	var s2 IntSlice
	var s3 SliceOfInt

	s1 = []int{0, 2, 4, 6, 8}
	s2 = s1
	s3 = s1

	s2 = IntSlice{9, 7, 5, 3, 1}
	s1 = s2
	s3 = s1

	s3 = SliceOfInt{0, -2, -4, -6, -8}
	s1 = s3
	s2 = s1
}