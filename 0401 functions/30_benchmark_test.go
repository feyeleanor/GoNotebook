package main
import "testing"

func Factorial(n int) (r int) {
  switch {
  case n < 0:
    panic(n)
  case n == 0:
    r = 1
  default:
    r = 1
    for ; n > 0; n-- {
      r *= n
    }
  }
  return
}

func SafeExecute(e func()) (r func(func())) {
  return func(f func()) {
    defer e()
    f()
  }
}

func InsideLoop(n, i int) {
  var errors int
  for ; n > 0; n-- {
  	for j := i; j > 0; j-- {
      SafeExecute(func() {
        if p := recover(); p != nil {
          errors++
        }
      })(func() { Factorial(5) })
    }
  }
}

func OutsideLoop(n, i int) {
  var errors int
  Execute := SafeExecute(func() {
    if p := recover(); p != nil {
      errors++
    }
  })
  for ; n > 0; n-- {
  	for j := i; j > 0; j-- {
      Execute(func() { Factorial(5) })
    }
  }
}

func BenchmarkInsideLoop1(b *testing.B) { InsideLoop(b.N, 1) }
func BenchmarkInsideLoop10(b *testing.B) { InsideLoop(b.N, 10) }
func BenchmarkInsideLoop100(b *testing.B) { InsideLoop(b.N, 100) }
func BenchmarkInsideLoop1000(b *testing.B) { InsideLoop(b.N, 1000) }

func BenchmarkOutsideLoop1(b *testing.B) { OutsideLoop(b.N, 1) }
func BenchmarkOutsideLoop10(b *testing.B) { OutsideLoop(b.N, 10) }
func BenchmarkOutsideLoop100(b *testing.B) { OutsideLoop(b.N, 100) }
func BenchmarkOutsideLoop1000(b *testing.B) { OutsideLoop(b.N, 1000) }
