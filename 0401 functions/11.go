package main
import . "fmt"

func main() {
  for i := 0; i < 10; i++ {
    Println("x =", FunctionA(i))
  }
}

func FunctionA(x int) {
  return x
}

func FunctionB(x int) int {}
