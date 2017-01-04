package main
import . "fmt"

type IntMap map[int] int
type MapOfInt map[int] int

func main() {
	m1 := IntMap{0: 9, 1: 7, 2: 5, 3: 3, 4: 1}
	m2 := MapOfInt{0: 0, 1: -2, 2: -4, 3: -6, 4: -8}
	m3 := map[int] int{0: 0, 1: 2, 2: 4, 3: 6, 4: 8}

	print_elements(m1)
	print_elements(m2)
	print_elements(m3)
}

func print_elements(m interface{}) {
	switch m := m.(type) {
	case map[int] int:
		for i, v := range m {
			Printf("%v: %v\n", i, v)
		}
	case IntMap:
		print_elements(as_unnamed_map(m))
	case MapOfInt:
		print_elements(as_unnamed_map(m))
	}
}

func as_unnamed_map(m map[int] int) map[int] int {
	return m
}