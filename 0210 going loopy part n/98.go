package main
import . "fmt"

type Iterable interface {
  Each(f interface{})
}

type IterableSlice []int
func (i *IterableSlice) Each(f interface{}) {
  switch f := f.(type) {
  case func(interface{}):
    for _, v := range *i {
      f(v)
    }
  case func(int, interface{}):
    for n, v := range *i {
      f(n, v)
    }
  }
}

func main() {
  print_values(&IterableSlice{ 0, 2, 4, 6, 8 })
}

func print_values(s Iterable) {
  var i int
  s.Each(func(v interface{}) {
    Printf("%v: %v\n", i, v)
    i++
  })
  s.Each(func(i int, v interface{}) {
    Printf("%v: %v\n", i, v)
  })
}
