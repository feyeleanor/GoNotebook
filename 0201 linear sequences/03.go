package main
import . "fmt"

func main() {
  i := 37
  for i, s := 0, []int{0, 2, 4, 6, 8}; i < len(s); i++ {
    Printf("%v: %v\n", i, s[i])
  }
  Println("i in outer scope:", i)
}
