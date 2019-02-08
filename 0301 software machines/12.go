package main
import "fmt"

type stack struct {
  int
  *stack
}

func (s stack) Push(v int) (r stack) {
  r = stack{v, &s}
  return
}

func (s stack) Pop() (v int, r stack) {
  return s.int, *s.stack
}

func (s stack) Depth() (r int) {
  for t := s.stack; t != nil; t = t.stack {
    r++
  }
  return
}

func (s *stack) PrintSum() {
  fmt.Printf("%d items: sum = %d\n", s.Depth(), s.Sum())
}

func (s stack) Sum() (r int) {
  for t, n := s, 0; t.stack != nil; r += n {
    n, t = t.Pop()
  }
  return
}

func main() {
  s1 := stack{}.Push(7)
  s2 := s1.Push(7).Push(11)

  s1 = s1.Push(2).Push(9).Push(4)
  s3 := s1.Push(17)

  s1 = s1.Push(3)

  s1.PrintSum()
  s2.PrintSum()
  s3.PrintSum()
}
