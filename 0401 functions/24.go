package main

import . "fmt"
import . "strconv"
import . "os"

func main() {
  var errors int

  for _, v := range Args[1:] {
    errors += SafeExecute(v,
      func(x int) {
        Printf("%v! = %v\n", x, Factorial(x))
      },
      func(p interface{}) {
        Printf("%v! is undefined\n", p)
      },
    )
  }

  Exit(errors)
}

func SafeExecute(v string, f func(int), e func(interface{})) (r int) {
  defer func() {
    if p := recover(); p != nil {
      r++
      e(p)
    }
  }()

  if x, e := Atoi(v); e == nil {
    f(x)
  } else {
    panic(v)
  }
  return
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
