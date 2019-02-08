package main
import . "fmt"

type IterableSlice []int

func main() {
  s := IterableSlice{ 0, 2, 4, 6, 8 }
  for i, v := range s {
    Printf("%v: %v\n", i, v)
  }
}
