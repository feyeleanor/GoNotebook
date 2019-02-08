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
        var x uint64

        if x, e = ParseUint(s, 10, 64); e == nil {
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

func Factorial(n uint64) (r uint64, e error) {
  if n == 0 {
    r = 1
  } else {
    if r, e = Factorial(n - 1); e == nil {
      if q := r * n; (q / n) != r {
        e = errors.New("unsigned integer overflow")
      } else {
        r = q
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
