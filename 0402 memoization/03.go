package main

import "math/big"
import "errors"
import . "fmt"

func main() {
  a := []string{}
  for i := 0; i < 256; i++ {
    a = append(a, Sprintf("%v", i))
  }

  var w int
  TryEach(a, func(s string) (e error) {
    Executor(func(err error) {
      Printf("%v! %v\n", s, err)
      e = err
    })(func() (e error) {
      var x *big.Int

      if x, e = ParseBigInt(s, 10); e == nil {
        if x, e = Factorial(x); e == nil {
          if len(x.Bits()) == w {
            Printf("%v! ", s)
          } else {
            w = len(x.Bits())
            Printf("\n%v: %v! ", w, s)
          }
        }
      }
      return
    })
    return
  })
  Println()
}

func Executor(e func(error)) (r func(func() error)) {
  return func(f func() error) {
    if err := f(); err != nil {
      e(err)
    }
  }
}

func Factorial(n *big.Int) (r *big.Int, e error) {
  switch n.Sign() {
  case -1:
    e = errors.New("is undefined")
  case 0:
    r = big.NewInt(1)
  default:
    if r, e = Factorial(succ(n)); e == nil {
      r.Mul(n, r)
    }
  }
  return
}

func succ(n *big.Int) *big.Int {
  return new(big.Int).Sub(n, big.NewInt(1))
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

func ParseBigInt(s string, base int) (r *big.Int, e error) {
  var x big.Int
  if _, ok := x.SetString(s, base); !ok {
    e = errors.New(Sprintf("ParseBigInt: parsing %v: invalid syntax", s))
  } else {
    r = &x
  }
  return
}
