package main
import "testing"

func Execute(f func()) {
  f()
}

func Compose(f func()) func() {
  return func() {
    f()
  }
}

func JITExecute(n, i int) {
  for ; n > 0; n-- {
    for j := i; j > 0; j-- {
      Execute(func() {})
    }
  }
}

func JITCompose(n, i int) {
  for ; n > 0; n-- {
    for j := i; j > 0; j-- {
      Compose(func() {})()
    }
  }
}

func PreCompose(n, i int) {
  for ; n > 0; n-- {
    c := Compose(func() {})
    for j := i; j > 0; j-- {
      c()
    }
  }
}

func BenchmarkExecute1(b *testing.B) { JITExecute(b.N, 1) }
func BenchmarkExecute10(b *testing.B) { JITExecute(b.N, 10) }
func BenchmarkExecute100(b *testing.B) { JITExecute(b.N, 100) }

func BenchmarkJITCompose1(b *testing.B) { JITCompose(b.N, 1) }
func BenchmarkJITCompose10(b *testing.B) { JITCompose(b.N, 10) }
func BenchmarkJITCompose100(b *testing.B) { JITCompose(b.N, 100) }

func BenchmarkPreCompose1(b *testing.B) { PreCompose(b.N, 1) }
func BenchmarkPreCompose10(b *testing.B) { PreCompose(b.N, 10) }
func BenchmarkPreCompose100(b *testing.B) { PreCompose(b.N, 100) }
