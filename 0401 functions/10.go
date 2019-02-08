package main
import . "fmt"

func main() {
  for i := 0; i < 10; i++ {
    Println("x =", Function(i))
  }
}

func Function(x int) int {
  return x
}
