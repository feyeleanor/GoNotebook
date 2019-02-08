package main
import . "fmt"

func main() {
  s := map[int] int{2: 4, 4: 8, 6: 12, 8: 16}
  print_sparse_array(s)
}

func print_sparse_array(s map[int] int) {
  n := len(s)
  for i := 0; n > 0; i++ {
    Printf("%v: ", i)
    if v, ok := s[i]; ok {
      Printf("%v", v)
      n--
    } else {
      Printf("0")
    }
    Println()
  }
}
