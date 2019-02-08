package main

import . "fmt"

type Signal struct{}
type Sequence []<-chan Signal

func launch_all(n int, f func(int, chan<- Signal)) {
  up := make(chan Signal)
  for c := 0; c < n; c++ {
    f(c, up)
  }
  for ; n > 0; n-- {
    <-up
  }
  close(up)
}

func NewSequence(n int) (r Sequence) {
  r = make(Sequence, n)
  launch_all(n, func(i int, up chan<- Signal) {
    c := make(chan Signal)
    r[i] = c
    go func(c chan<- Signal, i int) {
      up <- Signal{}
      for l := i * 2; l > 0; l-- {
        c <- Signal{}
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
  s := NewSequence(5)
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
