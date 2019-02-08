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

func (v *VM) read_program() interface{} {
  return v.m[v.PC]
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

type Interpreter struct {
  VM
  S    *stack
  l, r int
}

func main() {
  p := new(Interpreter)
  p.Load(
    Label("start"),
    Primitive(func() { p.S = p.S.Push(13) }),
    Primitive(func() { p.S = p.S.Push(28) }),
    Primitive(func() {
      p.l, p.S = p.S.Pop()
      p.r, p.S = p.S.Pop()
      p.S = p.S.Push(p.l + p.r)
    }),
    Primitive(func() {
      fmt.Printf("%v + %v = %v\n", p.l, p.r, p.S.int)
    }),
    Primitive(func() { os.Exit(0) }),
    Label("end"),
  ).Run()
}
