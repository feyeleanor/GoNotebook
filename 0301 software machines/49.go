package main
import "fmt"
import "errors"
import "strings"

type NO_OPERANDS struct{}
type SIZED struct{ int }
type TARGET struct{ int }
type TRANSFER struct{ d, s int }
type VALUE struct{ int }

type Memory []int

type OP interface {
  Apply(c *Core)
}

type CALL TARGET
type RETURN NO_OPERANDS
type GOTO TARGET
type MOVE TARGET
type DEBUG NO_OPERANDS
type PUSH VALUE
type POP VALUE
type ALLOC SIZED
type STORE TRANSFER
type ADD TRANSFER

func (i CALL) Apply(c *Core) {
  c.CallStack = &CallStack{i.int, c.CallStack}
}

func (RETURN) Apply(c *Core) {
  c.CallStack = c.CallStack.CallStack
  c.PC++
}

func (i GOTO) Apply(c *Core) {
  c.PC = i.int
}

func (i MOVE) Apply(c *Core) {
  c.PC += i.int
}

func (DEBUG) Apply(c *Core) {
  fmt.Printf("core: %v ticks, PC[%v], DS%v, M%v\n",
    c.Ticks,
    c.PC,
    c.DataStack,
    c.Memory,
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
  c.Memory[i.d] += c.Memory[i.s]
  c.PC++
}

type Program []OP

type CallStack struct {
  PC int
  *CallStack
}

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

type Core struct {
  Ticks   int
  Running bool
  e       error
  OP
  Program
  *DataStack
  *CallStack
  Memory
}

func NewCore() *Core {
  return &Core{CallStack: new(CallStack)}
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
  c := NewCore()
  c.Run(Program{
    CALL{2},
    GOTO{-1},
    ALLOC{2},
    DEBUG{},
    STORE{0, 3},
    STORE{1, 5},
    DEBUG{},
    ADD{0, 1},
    DEBUG{},
    RETURN{},
  })
}
