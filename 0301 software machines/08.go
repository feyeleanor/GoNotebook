package main
import "fmt"

type stack struct {
  data int
  tail *stack
}

func (s stack) Push(v int) (r stack) {
  r = stack{data: v, tail: &s}
  return
}

func (s stack) Pop() (v int, r stack) {
  return s.data, *s.tail
}

func (s stack) Depth() (r int) {
  for t := s.tail; t != nil; t = t.tail {
    r++
  }
  return
}

func (s *stack) append(n int) {
  t := s
  for ; t.tail != nil; t = t.tail {
  }
  *t = stack{data: n, tail: new(stack)}
}

func main() {
  var l, r int
  var s stack

  s = s.Push(1).Push(3)
  fmt.Printf("depth = %d\n", s.Depth())
  s.append(20)
  fmt.Printf("depth = %d\n", s.Depth())

  l, s = s.Pop()
  r, s = s.Pop()
  fmt.Printf("%d + %d = %d\n", l, r, l+r)

  l, s = s.Pop()
  fmt.Printf("l = %d\n", l)
  fmt.Printf("depth = %d\n", s.Depth())

  s.append(5)
  l, s = s.Pop()
  fmt.Printf("l = %d\n", l)
  fmt.Printf("depth = %d\n", s.Depth())
}
