package main
import "fmt"

type stack []int

func (s *stack) Push(data int) {
  *s = append(*s, data)
}

func (s *stack) Pop() (r int) {
  sp := len(*s) - 1
  r = (*s)[sp]
  *s = (*s)[:sp]
  return
}

func (s stack) Depth() int {
  return len(s)
}

func main() {
  s := new(stack)
  s.Push(1)
  s.Push(3)
  fmt.Printf("depth = %d\n", s.Depth())
  l := s.Pop()
  r := s.Pop()
  fmt.Printf("%d + %d = %d\n", l, r, l+r)
  fmt.Printf("depth = %d\n", s.Depth())
}
