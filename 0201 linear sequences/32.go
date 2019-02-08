package main
import . "fmt"

type Array [5]int

func main() {
  defer func() {
    recover()
  }()
  s := Array{0, 2, 4, 6, 8}
  for i := 0; ; i++ {
    Printf("%v: %v\n", i, s[i])
  }
}
