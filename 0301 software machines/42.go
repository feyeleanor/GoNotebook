package main
import "fmt"
import "errors"

type Instruction []int

type CALL Instruction
type RETURN Instruction
type GOTO Instruction
type MOVE Instruction

type Program []interface{}

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
  OP      interface{}
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
        c.CS = &CallStack{c.PC, c.CS}
        c.PC = op[0]
      case RETURN:
        c.PC, c.CS = c.CS.PC, c.CS.CallStack
        c.PC++
      default:
        panic("invalid instruction")
      }
      fmt.Println(c)
    }
  }
  return
}

func main() {
  c := new(Core)
  c.Run(Program{
    CALL{2},
    GOTO{5},
    MOVE{2},
    RETURN{},
    MOVE{-1},
  })
  fmt.Printf("Clock Cycles: %v, PC: %v %v\n",
    c.Ticks,
    c.PC,
    c.e,
  )
}
