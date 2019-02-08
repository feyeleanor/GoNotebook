package main
import "fmt"
import "errors"

type Instruction []int

type Program []Instruction

const (
  GOTO = iota
  MOVE
  CALL
  RETURN
)

type CallStack struct {
  PC int
  *CallStack
}

type Core struct {
  Ticks   int
  Running bool
  PC      int
  CS      *CallStack
  e       error
  OP      Instruction
  Program
}

func (c *Core) String() string {
  return fmt.Sprintf("ticks: %v, PC: %v, running: %v",
    c.Ticks,
    c.PC,
    c.Running,
  )
}

func (c *Core) LoadInstruction() {
  c.OP = c.Program[c.PC]
}

func (c *Core) CheckOperands(n int) {
  if len(c.OP) < n+1 {
    panic(fmt.Sprintf("PC[%v] insufficient operands", c.PC))
  }
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
      switch c.LoadInstruction(); c.OP[0] {
      case GOTO:
        if c.CheckOperands(1); c.e == nil {
          c.PC = c.OP[1]
        } else {
          panic("goto without target")
        }
      case MOVE:
        if c.CheckOperands(1); c.e == nil {
          c.PC = c.PC + c.OP[1]
        } else {
          panic("run: move without offset")
        }
      case CALL:
        if c.CheckOperands(1); c.e == nil {
          c.CS = &CallStack{c.PC, c.CS}
          c.PC = c.OP[1]
        }
      case RETURN:
        c.PC, c.CS = c.CS.PC, c.CS.CallStack
        c.PC++
      default:
        panic("unknown instruction")
      }
      fmt.Println(c)
    }
  }
  return
}

func main() {
  c := new(Core)
  c.Run(Program{
    {CALL, 2},
    {GOTO, 5},
    {MOVE, 2},
    {RETURN},
    {MOVE, -1},
  })
  fmt.Printf("Clock Cycles: %v, PC: %v %v\n",
    c.Ticks,
    c.PC,
    c.e,
  )
}
