package main

import . "fmt"
import . "strconv"
import . "os"

func main() {
  var errors int

  Executor := SafeExecutor(
    func() {
      if p := recover(); p != nil {
        Printf("%v! is undefined\n", p)
        errors++
      }
    })

  Range(Args[1:], func(s string) {
    Executor(
      func() {
        if x, e := Atoi(s); e == nil {
          Printf("%v! = %v\n", x, Factorial(x))
        } else {
          panic(s)
        }
      },
    )
  })

  Exit(errors)
}

func SafeExecutor(e func()) (r func(func())) {
  return func(f func()) {
    defer e()
    f()
  }
}

func Factorial(n int) (r int) {
  switch {
  case n < 0:
    panic(n)
  case n == 0:
    r = 1
  default:
    r = n * Factorial(n - 1)
  }
  return
}

func Range(a []string, f func(string)) {
  if len(a) > 0 {
    f(a[0])
    Range(a[1:], f)
  }
}
