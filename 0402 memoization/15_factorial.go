package main

import "math/big"
import "errors"
import . "fmt"

type Transform func(*big.Int) (*big.Int, error)

func FactorialCalculator(c Cache) (t Transform) {
  t = func(n *big.Int) (r *big.Int, e error) {
    switch r, e = c.Fetch(n); {
    case n.Sign() == -1:
      e = errors.New("is undefined")
    case e != nil:
    case r == nil:
      if r, e = t(succ(n)); e == nil {
        e = c.Store(n, r.Mul(n, r))
      }
    }
    return
  }
  return
}

func CachedFactorials(c Cache, f func(Transform) int) int {
  return c.
    Prefill(
      big.NewInt(0), big.NewInt(1)).
    Load().
    UseFor(func(c Cache) int {
      return f(
        FactorialCalculator(c))
    })
}

func (t Transform) Calculator(n string, f func(*big.Int, *big.Int) error) (r ErrorSource) {
  if x, e := ParseBigInt(n, 10); e == nil {
    n := CopyBigInt(x)
    r = func() (e error) {
      if x, e = t(CopyBigInt(n)); e == nil {
        e = f(CopyBigInt(n), x)
      }
      return
    }
  } else {
    r = func() error {
      return e
    }
  }
  return
}

func (t Transform) Execute(s string) error {
  return OnError(
    func(e error) error {
      Printf("%v! %v\n", s, e)
      return e
    },
  )(t.Calculator(s,
    func(n, x *big.Int) (e error) {
      Printf("%v! = %v\n", n, x)
      return
    },
  ))
}

func CalculateFactorials(args ...string) func(Transform) int {
  return func(t Transform) int {
    return TryEach(args, func(n string) error {
      return t.Execute(n)
    })
  }
}
