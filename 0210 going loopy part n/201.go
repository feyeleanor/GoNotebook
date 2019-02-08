package main

import "fmt"

func duplicate(v interface{}) interface{} {
  return v
}

type A struct { a, b int }

func main() {
  a := A{1, 2}
  fmt.Printf("a: %v\n", a)

  b := duplicate(a).(A)
  fmt.Printf("b: %v\n", b)
  b.a = 0
  fmt.Printf("b: %v\n", b)

  c := duplicate(&a).(*A)
  fmt.Printf("c: %v\n", c)
  c.b = 3
  fmt.Printf("a: %v\n", a)
  fmt.Printf("c: %v\n", c)
}
