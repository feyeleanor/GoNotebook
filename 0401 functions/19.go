package main

import . "fmt"
import . "strconv"
import . "os"

func main() {
  for _, v := range Args[1:] {
    CatchBadArguments(v, func(x int) {
      Printf("%v! = %v\n", x, Factorial(x))
    })
  }
}

func CatchBadArguments(v string, f func(int)) {
  defer func() {
    if p := recover(); p != nil {
      Printf("%v! is undefined\n", p)
    }
  }()

  if x, e := Atoi(v); e == nil {
    f(x)
  } else {
      Printf("%v! is undefined\n", v)
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
