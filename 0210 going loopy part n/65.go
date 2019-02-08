package main

import . "fmt"

func main() {
  c := MakeSequence(5)
  c.each(func(i, v int) {
    Printf("%v: %v\n", i, v)
  })
}

type Index int
type Sequence []chan Index

func MakeSequence(n int) (r Sequence) {
  r = make(Sequence, n)
  for i := 0; i < n; i++ {
    c := make(chan Index)
    r[i] = c
    go func(c chan Index, i int) {
      for l := i * 2; l > 0; l-- {
        c <- Index(i)
      }
      close(c)
    }(c, i)
  }
  return
}

func (s Sequence) each(f func(int, int)) {
  r := make([]int, len(s))
  for n := len(s); n > 0; {
    f := func(i Index, open bool) {
      if open {
        r[i]++
        Printf("%v", i)
      } else {
        n--
        s[i] = nil
      }
    }
    select {
    case _, ok := <-s[0]:
      f(0, ok)
    case _, ok := <-s[1]:
      f(1, ok)
    case _, ok := <-s[2]:
      f(2, ok)
    case _, ok := <-s[3]:
      f(3, ok)
    case _, ok := <-s[4]:
      f(4, ok)
    }
  }
  Println()
  for i, v := range r {
    f(i, v)
  }
}
