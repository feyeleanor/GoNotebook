package main

import "errors"
import . "fmt"
import . "strconv"
import . "os"

func main() {
  Exit(
    TryEach(Args[1:], func(s string) (e error) {
      Executor(func(err error) {
        Printf("%v! %v\n", s, err)
        e = err
      })(func() (e error) {
        var x int
        if x, e = Atoi(s); e == nil {
          if x, e = Factorial(x); e == nil {
            Printf("%v! = %v\n", s, x)
          }
        }
        return
      })
      return
    }),
  )
}

func Executor(e func(error)) (r func(func() error)) {
  return func(f func() error) {
    if err := f(); err != nil {
      e(err)
    }
  }
}

func Factorial(n int) (r int, e error) {
  switch {
  case n < 0:
    e = errors.New("is undefined")
  case n == 0:
    r = 1
  default:
    if r, e = Factorial(n - 1); e == nil {
      if r *= n; r < 0 {
        e = errors.New("integer overflow")
      }
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
