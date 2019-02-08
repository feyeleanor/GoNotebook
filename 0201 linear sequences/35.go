package main
import . "fmt"

func main() {
  print_array([4]int{0, 2, 4, 6})
  print_array([5]int{0, 2, 4, 6, 8})
  print_array([6]int{0, 2, 4, 6, 8, 10})
}

func print_array(s interface{}) {
  switch s := s.(type) {
  case [4]int:
    for i, v := range s {
      Printf("%v: %v\n", i, v)
    }
  case [5]int:
    for i, v := range s {
      Printf("%v: %v\n", i, v)
    }
  case [6]int:
    for i, v := range s {
      Printf("%v: %v\n", i, v)
    }
  }
}
