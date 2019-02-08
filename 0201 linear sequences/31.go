package main
import . "fmt"

func main() {
  defer func() {
    recover()
  }()
  s := [5]int{0, 2, 4, 6, 8}
  var i int
  for {
    Printf("%v: %v\n", i, s[i])
    i++
  }
}
