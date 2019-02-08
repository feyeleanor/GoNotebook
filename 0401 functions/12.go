package main
import . "fmt"

func main() {
  for i := 0; i < 10; i++ {
    var x, y int

    x, _ = Function(i)
    Printf("Function(%v) -> (%v, %v), ", i, x, y)

    _, y = Function(i)
    Printf("Function(%v) -> (%v, %v)\n", i, x, y)
  }
}

func Function(x int) (int, int) {
  return x, x * 2
}
