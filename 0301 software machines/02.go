package main
import "fmt"

type stack_status int

const (
  STACK_OK = stack_status(iota)
  STACK_OVERFLOW
  STACK_UNDERFLOW
)

type stack struct {
  data []int
}

func (s *stack) Push(data int) {
  s.data = append(s.data, data)
}

func (s *stack) Pop() (int, stack_status) {
  if s == nil || len(s.data) < 1 {
    return 0, STACK_UNDERFLOW
  }
  sp := len(s.data) - 1
  r := s.data[sp]
  s.data = s.data[:sp]
  return r, STACK_OK
}

func (s *stack) Depth() int {
  return len(s.data)
}

func main() {
  s := new(stack)
  s.Push(1)
  s.Push(3)
  fmt.Printf("depth = %d\n", s.Depth())
  l, _ := s.Pop()
  r, _ := s.Pop()
  fmt.Printf("%d + %d = %d\n", l, r, l+r)
  fmt.Printf("depth = %d\n", s.Depth())
}
