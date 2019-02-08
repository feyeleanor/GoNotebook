package main
import "os"
import "strconv"

var y int
func accumulate(x int) int {
  y += x
  return y
}

func main() {
  for _, v := range os.Args[1:] {
    x, _ := strconv.Atoi(v)
    accumulate(x)
  }
  os.Exit(accumulate(0))
}
