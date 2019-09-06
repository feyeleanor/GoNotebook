package main

import "math/big"
import "errors"
import . "fmt"

func ParseBigInt(s string, base int) (r *big.Int, e error) {
  var x big.Int
  if _, ok := x.SetString(s, base); !ok {
    e = errors.New(Sprintf("ParseBigInt: parsing %v: invalid syntax", s))
  } else {
    r = &x
  }
  return
}

func CopyBigInt(i *big.Int) *big.Int {
  return new(big.Int).Set(i)
}

func encode(v *big.Int) string {
  return v.Text(big.MaxBase)
}

func decode(s string) (*big.Int, error) {
  return ParseBigInt(s, big.MaxBase)
}

func Range(i *big.Int, f func(*big.Int) error) bool {
  return f(i) != nil || Range(i.Add(i, big.NewInt(1)), f)
}

type ErrorSource func() error
type ErrorPipe func(error) error

func OnError(p ErrorPipe) (r func(ErrorSource) error) {
  return func(s ErrorSource) (e error) {
    if e = s(); e != nil {
      e = p(e)
    }
    return
  }
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

func succ(n *big.Int) *big.Int {
  return new(big.Int).Sub(n, big.NewInt(1))
}
