package main
import . "fmt"

type IterableSlice []int
func (i IterableSlice) Each(f func(int, interface{})) {
  for n, v := range i {
    f(n, v)
  }
}

func main() {
  IterableSlice{ 0, 2, 4, 6, 8 }.Each(func(i int, v interface{}) {
    Printf("%v: %v\n", i, v)
  })
}
