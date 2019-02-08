package main
import . "fmt"

func main() {
  var i int
  s := []int{0, 2, 4, 6, 8}
  for _, _ = range s {
    Printf("%v: %v\n", i, s[i])
    i++
  }
}
