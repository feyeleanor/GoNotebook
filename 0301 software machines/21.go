package main
import "fmt"

func main() {
  var program = []interface{}{
    PUSH, 13,
    PUSH, 28,
    ADD,
    PRINT,
    EXIT,
  }
  interpret(program)
}

type stack struct {
  int
  *stack
}

func (s *stack) Push(v int) *stack {
  return &stack{v, s}
}

func (s *stack) Pop() (int, *stack) {
  return s.int, s.stack
}

type OPCODE int

const (
  PUSH = OPCODE(iota)
  ADD
  PRINT
  EXIT
)

func interpret(p []interface{}) {
  var l, r int
  S := new(stack)

  for PC := 0; ; PC++ {
    if op, ok := p[PC].(OPCODE); ok {
      switch op {
      case PUSH:
        PC++
        S = S.Push(p[PC].(int))
      case ADD:
        l, S = S.Pop()
        r, S = S.Pop()
        S = S.Push(l + r)
      case PRINT:
        fmt.Printf("%v + %v = %v\n", l, r, S.int)
      case EXIT:
        return
      }
    } else {
      return
    }
  }
}
