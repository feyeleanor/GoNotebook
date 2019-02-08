package main
import . "fmt"

func main() {
  s := func(c chan int) {
    for _, v := range []int{0, 2, 4, 6, 8} {
      c <- v
    }
    close(c)
  }
  f := func(c chan int) {
    go s(c)
    Printf("elements: %v\n", print_channel(c))
  }
  f(make(chan int))
  f(make(chan int, 16))
}

func print_channel(c chan int) (i int) {
  for v := range c {
    Printf("%v: %v\n", i, v)
    i++
  }
  return
}
