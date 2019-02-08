package main
import . "fmt"

func main() {
  c := MakeSequence(5)
  Printf("elements: %v\n", c.each(func(i, v int) {
    Printf("%v: %v\n", i, v)
  }))
}

type Sequence chan int

func (s Sequence) each(f func(int, int)) (i int) {
  for v := range s {
    f(i, v)
    i++
  }
  return
}

func MakeSequence(n int) (r Sequence) {
  r = make(Sequence, n)
  go func() {
    for i := 0; i < n; i++ {
      r <- i * 2
    }
    close(r)
  }()
  return r
}
