package main
import "fmt"
import "errors"
import "strings"

type Memory []int

type OP interface {
  Apply(c *Core)
}

type CALL int
type RETURN struct{}
type GOTO int
type MOVE int
type DEBUG struct{}
type PUSH int
type POP int
type ADD struct{}
type ALLOC int

func (i CALL) Apply(c *Core) {
  c.CallStack = &CallStack{int(i), c.CallStack}
}

func (RETURN) Apply(c *Core) {
  c.CallStack = c.CallStack.CallStack
  c.PC++
}

func (i GOTO) Apply(c *Core) {
  c.PC = int(i)
}

func (i MOVE) Apply(c *Core) {
  c.PC += int(i)
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
  c.DataStack = &DataStack{int(i), c.DataStack}
  c.PC++
}

func (i POP) Apply(c *Core) {
  c.Memory[int(i)] = c.int
  c.DataStack = c.DataStack.DataStack
  c.PC++
}

func (i ALLOC) Apply(c *Core) {
  c.Memory = make(Memory, int(i), int(i))
  c.PC++
}

func (ADD) Apply(c *Core) {
  r := c.DataStack.int
  c.DataStack = c.DataStack.DataStack
  c.int += r
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
  *CallStack
  *DataStack
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
    CALL(2),
    GOTO(-1),
    ALLOC(1),
    DEBUG{},
    PUSH(3),
    PUSH(5),
    DEBUG{},
    ADD{},
    DEBUG{},
    POP(0),
    DEBUG{},
    RETURN{},
  })
}
