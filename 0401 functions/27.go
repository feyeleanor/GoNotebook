package main

import . "fmt"
import . "strconv"
import . "os"

func main() {
  var errors int

  const ErrorHandler = func() {
    if p := recover(); p != nil {
      Printf("%v! is undefined\n", p)
      errors++
    }
  }

  for _, v := range Args[1:] {
    SafeExecute(
      func() {
        if x, e := Atoi(v); e == nil {
          Printf("%v! = %v\n", x, Factorial(x))
        } else {
          panic(v)
        }
      },
      ErrorHandler,
    )
  }

  Exit(errors)
}

func SafeExecute(f, e func()) {
  defer e()
  f()
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
