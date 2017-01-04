package main
import . "fmt"

func main() {
	s := []int{9, 7, 5, 3, 1}
	for i, s := 0, []int{0, 2, 4, 6, 8}; i < len(s); i++ {
		Printf("%v: %v\n", i, s[i])
	}
	for i := 0; i < len(s); i++ {
		Printf("%v: %v\n", i, s[i])
	}
}