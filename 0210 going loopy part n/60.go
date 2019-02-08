package main
import . "fmt"

func main() {
  c := make(chan int)
  go func() {
    for _, v := range []int{0, 2, 4, 6, 8} {
      c <- v
    }
    close(c)
  }()
  Printf("elements: %v\n", print_channel(c))
}

func print_channel(c chan int) (i int) {
  for v := range c {
    Printf("%v: %v\n", i, v)
    i++
  }
  return
}
