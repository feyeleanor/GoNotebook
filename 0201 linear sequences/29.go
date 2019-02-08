package main
import . "fmt"

func main() {
  s := [5]int{0, 2, 4, 6, 8}
  for i := 0; i < 5; i++ {
    Printf("%v: %v\n", i, s[i])
  }
}
