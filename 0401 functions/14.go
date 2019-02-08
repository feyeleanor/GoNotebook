package main

import . "fmt"
import "strconv"
import "os"

func main() {
  defer func() {
    if p := recover(); p != nil {
      Printf("%v! is undefined\n", p)
      os.Exit(1)
    }
  }()

  for _, v := range os.Args[1:] {
    if x, e := strconv.Atoi(v); e == nil {
      Printf("%v! = %v\n", x, Factorial(x))
    } else {
      panic(v)
    }
  }
}

func Factorial(n int) (r int) {
  switch {
  case n < 0:
    panic(n)
  case n == 0:
    r = 1
  default:
    r = 1
    for ; n > 0; n-- {
      r *= n
    }
  }
  return
}
