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

type RegMachine struct {
  R []int
  VM
}

func (r *RegMachine) Clear() {
  r.R = make([]int, 2, 2)
}

func (r *RegMachine) read_value() int {
  r.PC++
  return r.read_program().(int)
}

func (r *RegMachine) LoadValue() {
  r.R[r.read_value()] = r.read_value()
}

func (r *RegMachine) Add() {
  i := r.read_value()
  j := r.read_value()
  r.R[i] += r.R[j]
}

func (r *RegMachine) JumpIfNotZero() {
  if r.R[r.read_value()] != 0 {
    r.PC = r.read_value()
  } else {
    r.PC++
  }
}

func main() {
  r := new(RegMachine)
  print_state := Primitive(func() {
    fmt.Printf("pc[%v]: %v, Acc: %v\n",
      r.PC,
      r.m[r.PC],
      r.R,
    )
  })
  r.Load(
    Primitive(r.Clear),
    Primitive(r.LoadValue), 0, 13,
    Primitive(r.LoadValue), 1, -1,
    Label("dec"),
    Primitive(r.Add), 0, 1,
    print_state,
    Primitive(r.JumpIfNotZero), 0, Label("dec"),
    print_state,
    Primitive(func() {
      os.Exit(0)
    }),
  ).Run()
}
