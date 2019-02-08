package main
import "testing"

func FactorialA(n int) (r int) {
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

func FactorialB(n int) (r int) {
  if n < 0 {
    panic(n)
  }
  for r = 1; n > 0; n-- {
    r *= n
  }
  return
}

func BenchmarkFactorialA0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FactorialA(0)
	}
}

func BenchmarkFactorialA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FactorialA(1)
	}
}

func BenchmarkFactorialA20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FactorialA(20)
	}
}

func BenchmarkFactorialB0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FactorialB(0)
	}
}

func BenchmarkFactorialB1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FactorialB(1)
	}
}

func BenchmarkFactorialB20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FactorialB(20)
	}
}
