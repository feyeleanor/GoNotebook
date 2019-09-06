package main

import "math/big"
import "bufio"
import "errors"
import . "fmt"
import "io"
import "os"

const FILE = "13.txt"

var cache map[string] string

func init() {
  cache = map[string] string{ "0": "1" }

  f, _ := os.Open(FILE)
  defer f.Close()

  r := bufio.NewReader(f)
  Range(big.NewInt(1), func(i *big.Int) (e error) {
    s := i.Text(big.MaxBase)
    switch cache[s], e = r.ReadString('\n'); {
    case e == io.EOF && len(cache[s]) > 0:
      fallthrough
    case e == nil:
      x := cache[s]
      cache[s] = x[:len(x) - 1]
    default:
      delete(cache, s)
    }
    return
  })
}

func Range(i *big.Int, f func(*big.Int) error) {
  if f(i) != nil {
    return
  }
  Range(i.Add(i, big.NewInt(1)), f)
}

func writeCache(f *bufio.Writer, n, v *big.Int) (e error) {
  x := v.Text(big.MaxBase)
  cache[n.Text(big.MaxBase)] = x
  _, e = f.WriteString(x + "\n")
  return
}

func usingCache(s string, f func(*bufio.Writer) int) int {
  o, _ := os.OpenFile(s, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0755)
  w := bufio.NewWriter(o)

  defer func() {
    w.Flush()
    o.Close()
  }()

  return f(w)
}

func main() {
  os.Exit(
    usingCache(FILE, func(w *bufio.Writer) int {
      F := Factorial(w)
      return TryEach(os.Args[1:], func(s string) (e error) {
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
      })
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

func Factorial(w *bufio.Writer) (f func(*big.Int) (*big.Int, error)) {
  f = func(n *big.Int) (r *big.Int, e error) {
    if q, ok := cache[n.Text(big.MaxBase)]; ok {
      r, e = ParseBigInt(q, big.MaxBase)
    } else {
      if n.Sign() == -1 {
        e = errors.New("is undefined")
      } else {
        if r, e = f(succ(n)); e == nil {
          e = writeCache(w, n, r.Mul(n, r))
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
