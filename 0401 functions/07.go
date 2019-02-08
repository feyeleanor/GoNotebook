package main
import "os"
import "strconv"

func MakeAccumulator() func(int) int {
  var sum int

  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  accumulate := MakeAccumulator()
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    accumulate(x)
  }
  os.Exit(accumulate(0))
}
