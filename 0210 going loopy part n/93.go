package main
import . "fmt"

type IterableSlice []int

func main() {
  print_values(IterableSlice{ 0, 2, 4, 6, 8 })
}

func print_values(s []int) {
  for i, v := range s {
    Printf("%v: %v\n", i, v)
  }
}
