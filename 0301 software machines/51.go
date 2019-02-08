package main
import "fmt"
import "errors"
import "strings"

type PC int
type NO_OPERANDS struct{}
type SIZED struct{ int }
type JUMP_TARGET struct{ PC }
type TRANSFER struct{ d, s int }
type VALUE struct{ int }

type R [2]int
type Memory []int

type OP interface {
  Apply(c *Core)
}

type CALL JUMP_TARGET
type RETURN NO_OPERANDS
type GOTO JUMP_TARGET
type MOVE JUMP_TARGET
type DEBUG NO_OPERANDS
type PUSH VALUE
type POP VALUE
type ALLOC SIZED
type STORE TRANSFER
type ADD TRANSFER
type LOAD TRANSFER

func (i CALL) Apply(c *Core) {
  c.Registers.Registers = &Registers{
    R:         c.R,
    Memory:    c.Memory,
    OP:        c.OP,
    PC:        c.PC,
    DataStack: c.DataStack,
    Registers: c.Registers.Registers,
  }
  c.PC = i.PC
}

func (RETURN) Apply(c *Core) {
  c.R = c.Registers.Registers.R
  c.Memory = c.Registers.Registers.Memory
  c.OP = c.Registers.Registers.OP
  c.PC = c.Registers.Registers.PC
  c.DataStack = c.Registers.Registers.DataStack
  c.Registers.Registers = c.Registers.Registers.Registers
  c.PC++
}

func (i GOTO) Apply(c *Core) {
  c.PC = i.PC
}

func (i MOVE) Apply(c *Core) {
  c.PC += i.PC
}

func (DEBUG) Apply(c *Core) {
  fmt.Printf("core: %v ticks, PC[%v], DS%v, M%v, R%v\n",
    c.Ticks,
    c.PC,
    c.DataStack,
    c.Memory,
    c.R,
  )
  c.PC++
}

func (i PUSH) Apply(c *Core) {
  c.DataStack = &DataStack{i.int, c.DataStack}
  c.PC++
}

func (i POP) Apply(c *Core) {
  c.Memory[i.int] = c.int
  c.DataStack = c.DataStack.DataStack
  c.PC++
}

func (i ALLOC) Apply(c *Core) {
  c.Memory = make(Memory, i.int, i.int)
  c.DataStack = &DataStack{0, c.DataStack}
  c.PC++
}

func (i STORE) Apply(c *Core) {
  c.Memory[i.d] = i.s
  c.PC++
}

func (i ADD) Apply(c *Core) {
  c.R[i.d] += c.R[i.s]
  c.PC++
}

func (i LOAD) Apply(c *Core) {
  c.R[i.d] = c.Memory[i.s]
  c.PC++
}

type Program []OP

type DataStack struct {
  int
  *DataStack
}

func (d *DataStack) String() string {
  t := []string{}
  for ; d != nil; d = d.DataStack {
    t = append(t, fmt.Sprint(d.int))
  }
  return "[" + strings.Join(t, " ") + "]"
}

type Registers struct {
  R
  Memory
  OP
  PC
  *DataStack
  *Registers
}

type Core struct {
  Ticks   int
  Running bool
  e       error
  Program
  Registers
}

func (c *Core) LoadInstruction() {
  c.OP = c.Program[c.PC]
}

func (c *Core) Run(p Program) (e error) {
  defer func() {
    if x := recover(); x != nil {
      c.e = errors.New(fmt.Sprintf("run: %v", x))
      c.Running = false
    }
  }()

  if c.Running {
    panic("processor busy")
  } else {
    c.Program = p
    for c.PC, c.Running = 0, true; c.Running; c.Ticks++ {
      c.LoadInstruction()
      c.Apply(c)
    }
  }
  return
}

func main() {
  c := new(Core)
  c.Run(Program{
    CALL{3},
    DEBUG{},
    GOTO{-1},
    ALLOC{2},
    DEBUG{},
    STORE{0, 3},
    STORE{1, 5},
    DEBUG{},
    LOAD{1, 0},
    DEBUG{},
    ADD{0, 1},
    DEBUG{},
    LOAD{1, 1},
    ADD{0, 1},
    DEBUG{},
    RETURN{},
  })
}
