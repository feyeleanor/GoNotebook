package main
import . "fmt"

func main() {
  c := sequence(5)
  Printf("elements: %v\n", print_sequence(c))
}

func sequence(n int) (r chan int) {
  r = make(chan int, n)
  go func() {
    for i := 0; i < n; i++ {
      r <- i * 2
    }
    close(r)
  }()
  return r
}

func print_sequence(c chan int) (i int) {
  for v := range c {
    Printf("%v: %v\n", i, v)
    i++
  }
  return
}
