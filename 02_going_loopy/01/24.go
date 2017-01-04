package main
import . "fmt"

func main() {
	m := map[int] int{ 0: 0, 1: 2, 2: 4, 3: 6, 4: 8 }
	for i, v := range m {
		Printf("%v: %v\n", i, v)
	}
}