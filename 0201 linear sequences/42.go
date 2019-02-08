package main
import . "fmt"

func main() {
  a := [...]int{0, 2, 4, 6, 8}
  print_array(a[:])
}

func print_array(s []int) {
  for i, v := range s {
    Printf("%v: %v\n", i, v)
  }
}
