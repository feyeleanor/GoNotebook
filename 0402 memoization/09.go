package main

import "math/big"
import "errors"
import . "fmt"
import . "os"

func main() {
  F := Factorial()
  Exit(
    TryEach(Args[1:], func(s string) (e error) {
      Executor(func(err error) {
        Printf("%v! %v\n", s, err)
        e = err
      })(func() (e error) {
        var x *big.Int

        if x, e = ParseBigInt(s, 10); e == nil {
          if x, e = F(x); e == nil {
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

func Factorial() (f func(*big.Int) (*big.Int, error)) {
  cache := make(map[int64]*big.Int)

  f = func(n *big.Int) (r *big.Int, e error) {
    i := n.Int64()
    if q, ok := cache[i]; ok {
      r = big.NewInt(0).Set(q)
    } else {
      switch n.Sign() {
      case -1:
        e = errors.New("is undefined")
      case 0:
        cache[i] = big.NewInt(1)
        r = big.NewInt(1)
      default:
        if r, e = f(succ(n)); e == nil {
          q = big.NewInt(0).Mul(n, r)
          cache[i] = q
          r = big.NewInt(0).Set(q)
        }
      }
    }
    return
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
