package main
import . "fmt"

func main() {
	for i, s := 0, []int{0, 2, 4, 6, 8}; i < len(s); i++ {
		Printf("%v: %v\n", i, s[i])
	}
}