package main
import . "fmt"

func main() {
  print_array([5]int{0, 2, 4, 6, 8})
}

func print_array(s [5]int) {
  for i, v := range s {
    Printf("%v: %v\n", i, v)
  }
}
