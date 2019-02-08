package main
import . "fmt"

type NumberSequence interface {
  Each(func(int, int)) int
}

type NumberSlice []int
func (i NumberSlice) Each(f func(int, int)) (n int) {
  var v int
  for n, v = range i {
    f(n, v)
  }
  n++
  return
}

func main() {
  print_values(NumberSlice{ 0, 2, 4, 6, 8 })
}

func print_values(s NumberSequence) {
  Printf("elements: %v\n", s.Each(func(i, v int) {
    Printf("%v: %v\n", i, v)
  }))
}
