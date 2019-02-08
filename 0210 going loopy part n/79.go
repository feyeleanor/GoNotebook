package main

import . "fmt"

type Signal struct{}
type Output chan<- Signal
type Sequence []<-chan Signal

func (o Output) Trigger() {
  o <- Signal{}
}

func (s Sequence) init(f func(int, Output)) {
  n := len(s)
  up := make(chan Signal)
  for c := 0; c < n; c++ {
    f(c, up)
  }
  for ; n > 0; n-- {
    <-up
  }
  close(up)
}

func MakeSequence(n int) (r Sequence) {
  r = make(Sequence, n)
  r.init(func(i int, up Output) {
    c := make(chan Signal)
    r[i] = c
    go func(c Output, i int) {
      up.Trigger()
      for l := i * 2; l > 0; l-- {
        c.Trigger()
      }
      close(c)
    }(c, i)
  })
  return
}

func (s Sequence) Next(f func(int, bool)) {
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

func main() {
  s := MakeSequence(5)
  r := make([]int, len(s))

  for n := len(s); n > 0; {
    s.Next(func(i int, open bool) {
      if open {
        r[i]++
        Printf("%v", i)
      } else {
        n--
        s[i] = nil
      }
    })
  }
  Println()
  for i, v := range r {
    Printf("%v: %v\n", i, v)
  }
}
