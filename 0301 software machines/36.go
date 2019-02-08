package main
import "fmt"
import "os"

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

type Primitive func()
type Label string
type labels map[Label]int

type VM struct {
  PC int
  m  []interface{}
  labels
}

func (v *VM) Load(program ...interface{}) *VM {
  v.labels = make(labels)
  v.PC = -1
  for _, token := range program {
    v.assemble(token)
  }
  return v
}

func (v *VM) assemble(token interface{}) {
  switch t := token.(type) {
  case Label:
    if i, ok := v.labels[t]; ok {
      v.m = append(v.m, i)
    } else {
      v.labels[t] = v.PC
    }
  default:
    v.m = append(v.m, token)
    v.PC++
  }
}

func (v *VM) Run() {
  v.PC = -1
  for {
    v.PC++
    v.read_program().(Primitive)()
  }
}

func (v *VM) read_program() interface{} {
  return v.m[v.PC]
}

type StackMachine struct {
  *stack
  VM
}

func (s *StackMachine) Push() {
  s.PC++
  s.stack = s.stack.Push(s.read_program().(int))
}

func (s *StackMachine) Add() {
  var l, r int
  l, s.stack = s.Pop()
  r, s.stack = s.Pop()
  s.stack = s.stack.Push(l + r)
}

func (s *StackMachine) JumpIfNotZero() {
  s.PC++
  if s.int != 0 {
    s.PC = s.m[s.PC].(int)
  }
}

func main() {
  s := new(StackMachine)
  print_state := Primitive(func() {
    fmt.Printf("pc[%v]: %v, TOS: %v\n",
      s.PC,
      s.m[s.PC],
      s.int,
    )
  })
  s.Load(
    Primitive(s.Push), 13,
    Label("dec"),
    Primitive(func() {
      s.stack = s.stack.Push(-1)
    }),
    Primitive(s.Add),
    print_state,
    Primitive(s.JumpIfNotZero), Label("dec"),
    print_state,
    Primitive(func() {
      os.Exit(0)
    }),
  ).Run()
}
