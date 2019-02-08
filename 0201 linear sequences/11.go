package main
import . "fmt"

func main() {
  s := []int{0, 2, 4, 6, 8}
  for i, _ := range s {
    Printf("%v: %v\n", i, s[i])
  }
}
