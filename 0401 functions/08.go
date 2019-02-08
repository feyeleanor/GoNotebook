package main
import "os"
import "strconv"

type Accumulator int

func (a *Accumulator) Add(y int) int {
  *a += Accumulator(y)
  return int(*a)
}

func main() {
  var a Accumulator
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    a.Add(x)
  }
  os.Exit(a.Add(0))
}
