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

type Core struct {
  Ticks   int
  Running bool
  PC      int
  CS      []int
  e       error
  OP      Instruction
}

func NewCore(CSSize int) Core {
  return Core{
    CS: make([]int, 0, CSSize),
  }
}

func (c *Core) String() string {
  return fmt.Sprintf("ticks: %v, PC: %v, running: %v", c.Ticks, c.PC, c.Running)
}

func (c *Core) LoadInstruction(p Program) {
  switch l := len(p); {
  case c.PC < 0 || c.PC >= l:
    c.e = errors.New("load: program counter invalid")
  default:
    c.OP = p[c.PC]
  }
}

func (c *Core) CheckOperands(n int) {
  if len(c.OP) < n+1 {
    c.e = errors.New(fmt.Sprintf("core: PC[%v] insufficient operands", c.PC))
  }
}

func (c *Core) Run(p Program) (e error) {
  if c.Running {
    e = errors.New("run: processor busy")
  } else {
    for c.PC, c.Running = 0, true; c.Running; c.Ticks++ {
      if c.LoadInstruction(p); c.e == nil {
        if len(c.OP) == 0 {
          c.e = errors.New("run: not an instruction")
        } else {
          switch c.OP[0] {
          case GOTO:
            if c.CheckOperands(1); c.e == nil {
              c.PC = c.OP[1]
            } else {
              c.e = errors.New("run: goto without target")
            }
          case MOVE:
            if c.CheckOperands(1); c.e == nil {
              c.PC = c.PC + c.OP[1]
            } else {
              c.e = errors.New("run: move without offset")
            }
          case CALL:
            if c.CheckOperands(1); c.e == nil {
              if top := len(c.CS); top < cap(c.CS)-1 {
                c.CS = c.CS[:top+1]
                c.CS[top] = c.PC
                c.PC = c.OP[1]
              } else {
                c.e = errors.New("run: call stack full")
              }
            }
          case RETURN:
            if top := len(c.CS); top > 0 {
              c.PC, c.CS = c.CS[top-1], c.CS[:top]
            } else {
              c.e = errors.New("run: call stack empty")
            }
            c.PC++
          default:
            c.e = errors.New("run: unknown instruction")
          }
        }
      }
      if c.e != nil {
        c.Running = false
        break
      }
      fmt.Println(c)
    }
  }
  return
}

func main() {
  c := NewCore(10)
  c.Run(Program{
    {CALL, 2},
    {GOTO, 5},
    {MOVE, 2},
    {RETURN},
    {MOVE, -1},
  })
  fmt.Printf("Clock Cycles: %v, PC: %v %v\n", c.Ticks, c.PC, c.e)
}
