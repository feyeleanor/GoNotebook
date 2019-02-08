package main
import "fmt"

func main() {
  a := NewAssembler("noop", "load", "store")
  p := Program{
    a.Op("noop"),
    a.Op("load", 1),
    a.Op("store", 1, 2),
    a.Op("invalid", 3, 4, 5),
  }
  p.Disassemble(a)

  for _, v := range p {
    if len(v.Operands()) == 2 {
      v.Execute(func(o []int) {
        o[0] += o[1]
      })
      println("op =", v.Opcode(), "result =", v.Operands()[0])
    }
  }
}

type Operation func(o []int)

type Executable interface {
  Opcode() int
  Operands() []int
  Execute(op Operation)
}

type Program []Executable

func (p Program) Disassemble(a Assembler) {
  for _, v := range p {
    fmt.Println(a.Disassemble(v))
  }
}

const INVALID_OPCODE = -1

type Instruction []int

func (i Instruction) Opcode() int {
  if len(i) == 0 {
    return INVALID_OPCODE
  }
  return i[0]
}

func (i Instruction) Operands() []int {
  if len(i) < 2 {
    return []int{}
  }
  return i[1:]
}

func (i Instruction) Execute(op Operation) {
  op(i.Operands())
}

type Assembler struct {
  opcodes map[string]int
  names   map[int]string
}

func NewAssembler(names ...string) (a Assembler) {
  a = Assembler{
    make(map[string]int),
    make(map[int]string),
  }
  a.Define(names...)
  return
}

func (a Assembler) Define(names ...string) {
  for _, name := range names {
    a.opcodes[name] = len(a.names)
    a.names[len(a.names)] = name
  }
}

func (a Assembler) Op(name string, params ...int) (i Instruction) {
  i = make(Instruction, len(params)+1)
  if opcode, ok := a.opcodes[name]; ok {
    i[0] = opcode
  } else {
    i[0] = INVALID_OPCODE
  }
  copy(i[1:], params)
  return
}

func (a Assembler) Disassemble(e Executable) (s string) {
  if name, ok := a.names[e.Opcode()]; ok {
    s = name
    if params := e.Operands(); len(params) > 0 {
      s = fmt.Sprintf("%v\t%v", s, params[0])
      for _, v := range params[1:] {
        s = fmt.Sprintf("%v, %v", s, v)
      }
    }
  } else {
    s = "unknown"
  }
  return
}
