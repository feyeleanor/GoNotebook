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
        var x *big.Int

        if x, e = ParseBigInt(s, 10); e == nil {
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

var cache []*big.Int

func Factorial(n *big.Int) (r *big.Int, e error) {
  if n.Cmp(big.NewInt(int64(len(cache)))) == -1 {
    r = cache[int(n.Int64())]
    Printf("cache[%v] = %v\n", n, r)
  } else {
    Printf("caching %v!\n", n)
    switch n.Sign() {
    case -1:
      e = errors.New("is undefined")
    case 0:
      r = big.NewInt(1)
      cache = append(cache, r)
    default:
      if r, e = Factorial(succ(n)); e == nil {
        r.Mul(n, r)
        cache = append(cache, r)
      }
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
