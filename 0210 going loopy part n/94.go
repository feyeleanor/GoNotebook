package main
import . "fmt"

type IterableSlice []int
func (i IterableSlice) Each(f func(interface{})) {
  for _, v := range i {
    f(v)
  }
}

func main() {
  i := 0
  IterableSlice{ 0, 2, 4, 6, 8 }.Each(func(v interface{}) {
    Printf("%v: %v\n", i, v)
    i++
  })
}
