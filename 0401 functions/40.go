package main

import "errors"
import . "fmt"
import . "strconv"
import . "os"

func main() {
  Exit(
    TryEach(Args[1:], func(s string) (e error) {
      SafeExecutor(func() {
        if p := recover(); p != nil {
          Printf("%v! %v\n", s, p)
          e = errors.New("")
        }
      })(func() {
        if x, e := Atoi(s); e == nil {
          Printf("%v! = %v\n", x, Factorial(x))
        } else {
          panic(s)
        }
      })
      return
    }),
  )
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
    panic("is undefined")
  case n == 0:
    r = 1
  default:
    if r = n * Factorial(n - 1); r < 0 {
      panic("integer overflow")
    }
  }
  return
}

func TryEach(a []string, f func(string) error) (r int) {
  switch {
  case len(a) == 0:
  case f(a[0]) != nil:
    r++
    fallthrough
  default:
    r += TryEach(a[1:], f)
  }
  return
}
