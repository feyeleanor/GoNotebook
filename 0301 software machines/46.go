package main
import "fmt"
import "errors"
import "strings"

type Instruction []int

type CALL Instruction
type RETURN Instruction
type GOTO Instruction
type MOVE Instruction
type DEBUG Instruction
type PUSH Instruction
type POP Instruction
type ADD Instruction

func (DEBUG) Report(c *Core) {
  fmt.Printf("core: %v ticks, PC[%v], DS%v\n",
    c.Ticks,
    c.PC,
    c.DataStack,
  )
}

type Program []interface{}

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
  OP      interface{}
  Program
  *CallStack
  *DataStack
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
      switch op := c.OP.(type) {
      case GOTO:
        c.PC = op[0]
      case MOVE:
        c.PC = c.PC + op[0]
      case CALL:
        c.CallStack = &CallStack{op[0], c.CallStack}
      case RETURN:
        c.CallStack = c.CallStack.CallStack
        c.PC++
      case DEBUG:
        op.Report(c)
        c.PC++
      case PUSH:
        c.DataStack = &DataStack{op[0], c.DataStack}
        c.PC++
      case POP:
        c.DataStack = c.DataStack.DataStack
        c.PC++
      case ADD:
        r := c.int
        c.DataStack = c.DataStack.DataStack
        c.int += r
        c.PC++
      default:
        panic("invalid instruction")
      }
    }
  }
  return
}

func main() {
  c := NewCore()
  c.Run(Program{
    PUSH{3},
    PUSH{5},
    DEBUG{},
    ADD{},
    DEBUG{},
    POP{},
    DEBUG{},
  })
}
