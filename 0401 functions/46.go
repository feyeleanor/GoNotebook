package main

import "math/big"
import "errors"
import . "fmt"
import . "os"

func main() {
  Exit(
    TryEach(Args[1:], func(s string) (e error) {
      Executor(func(err error) {
        Printf("%v! %v\n", s, err)
        e = err
      })(func() (e error) {
        var x *Uint

        if x, e = ParseBigUint(s, 10); e == nil {
          Printf("%v! = %v\n", s, Factorial(x))
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

func Factorial(n *Uint) (r *Uint) {
  if n.Sign() == 0 {
    r = NewUint(1)
  } else {
    r = NewUint(0).Mul(n, Factorial(succ(n)))
  }
  return
}

func succ(n *Uint) (r *Uint) {
  r = NewUint(0).Sub(n, NewUint(1))
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

type Uint struct {
  *big.Int
}

func NewUint(x uint64) (r *Uint) {
  return &Uint{ big.NewInt(int64(x)) }
}

func ParseBigUint(s string, base int) (r *Uint, e error) {
  x := NewUint(0)
  if _, ok := x.SetString(s, base); !ok || x.Sign() == -1 {
    e = errors.New(Sprintf("ParseBigUint: parsing %v: invalid syntax", s))
  } else {
    r = x
  }
  return
}

func (u *Uint) Sub(x, y *Uint) *Uint {
  if x.Cmp(y.Int) == -1 {
    panic("negative result for unsigned integer operation")
  } else {
    u.Int.Sub(x.Int, y.Int)
  }
  return u
}

func (u *Uint) Mul(x, y *Uint) *Uint {
  u.Int.Mul(x.Int, y.Int)
  return u
}
