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

type AccMachine struct {
  int
  VM
}

func (a *AccMachine) Clear() {
  a.int = 0
}

func (a *AccMachine) LoadValue() {
  a.PC++
  a.int = a.read_program().(int)
}

func (a *AccMachine) Add() {
  a.PC++
  a.int += a.read_program().(int)
}

func (a *AccMachine) JumpIfNotZero() {
  a.PC++
  if a.int != 0 {
    a.PC = a.m[a.PC].(int)
  }
}

func main() {
  a := new(AccMachine)
  print_state := Primitive(func() {
    fmt.Printf("pc[%v]: %v, Acc: %v\n",
      a.PC,
      a.m[a.PC],
      a.int,
    )
  })
  a.Load(
    Primitive(a.Clear),
    Primitive(a.LoadValue), 13,
    Label("dec"),
    Primitive(a.Add), -1,
    print_state,
    Primitive(a.JumpIfNotZero), Label("dec"),
    print_state,
    Primitive(func() {
      os.Exit(0)
    }),
  ).Run()
}
