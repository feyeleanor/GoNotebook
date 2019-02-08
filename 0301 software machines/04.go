package main
import "fmt"

type stack []int

func (s stack) Pop() (int, stack) {
  sp := len(s) - 1
  return s[sp], s[:sp]
}

func main() {
  s := make(stack, 0)
  s = append(s, 1, 3)
  fmt.Printf("depth = %d\n", len(s))
  var l, r int
  l, s = s.Pop()
  r, s = s.Pop()
  fmt.Printf("%d + %d = %d\n", l, r, l+r)
  fmt.Printf("depth = %d\n", len(s))
}
