package main
import . "fmt"

func main() {
  for i := 0; i < 10; i++ {
    Procedure(i)
  }
}

func Procedure(x int) {
  Println("x =", x)
}
