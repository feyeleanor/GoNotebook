package main

import "errors"
import . "fmt"
import . "strconv"
import . "os"

func main() {
  Exit(
    TryEach(Args[1:], func(s string) (e error) {
      SafeExecutor(func(err error) {
        Printf("%v! %v\n", s, err)
        e = err
      })(func() (e error) {
        var x int
        if x, e = Atoi(s); e == nil {
          Printf("%v! = %v\n", x, Factorial(x))
        }
        return
      })
      return
    }),
  )
}

func SafeExecutor(e func(error)) (r func(func() error)) {
  return func(f func() error) {
    var err error
    defer func() {
      switch p := recover(); {
      case p != nil:
        e(errors.New(Sprint(p)))
      case err != nil:
        e(err)
      }
    }()
    err = f()
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
