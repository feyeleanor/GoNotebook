package main

import "bufio"
import "fmt"
import "os"
import "testing"

const FILE LineFile = "20.txt"

type LineWriter struct {
  *bufio.Writer
}

func NewLineWriter(f *os.File) (r *LineWriter) {
  return &LineWriter{ bufio.NewWriter(f) }
}

func (l *LineWriter) WriteLine(s string) (int, error) {
  return l.Writer.WriteString(s + "\n")
}

type LineFile string

func (l LineFile) Append(f func(*LineWriter) error) (e error) {
  var o *os.File

  o, e = os.OpenFile(string(l), os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
  if e == nil {
    defer o.Close()
    w := NewLineWriter(o)
    defer w.Flush()
    e = f(w)
  }
  if e != nil {
    fmt.Println(l, "error writing", e.Error())
  }
  return
}

func (l LineFile) AppendLine(s string) (e error) {
  var o *os.File

  o, e = os.OpenFile(string(l), os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
  if e == nil {
    defer o.Close()
    w := bufio.NewWriter(o)
    defer w.Flush()
    w.WriteString(s + "\n")
  }
  if e != nil {
    fmt.Println(l, "error writing", e.Error())
  }
  return
}

func BenchmarkAppend0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FILE.Append(func(*LineWriter) (e error) {
      return
    })
	}
}

func BenchmarkAppend1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FILE.Append(func(w *LineWriter) (e error) {
      w.WriteLine("1")
      return
    })
	}
}

func BenchmarkAppendLine1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FILE.AppendLine("1")
	}
}

func BenchmarkAppend10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FILE.Append(func(w *LineWriter) (e error) {
      for j := 10; j > 0; j-- {
        w.WriteLine("1")
      }
      return
    })
	}
}

func BenchmarkAppendIndividually10(b *testing.B) {
	for i := 0; i < b.N; i++ {
    for j := 10; j > 0; j-- {
  		FILE.Append(func(w *LineWriter) (e error) {
        w.WriteLine("1")
        return
      })
    }
	}
}

func BenchmarkAppendLine10(b *testing.B) {
	for i := 0; i < b.N; i++ {
    for j := 10; j > 0; j-- {
  		FILE.AppendLine("1")
    }
	}
}

func BenchmarkAppend100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FILE.Append(func(w *LineWriter) (e error) {
      for j := 100; j > 0; j-- {
        w.WriteLine("1")
      }
      return
    })
	}
}

func BenchmarkAppend100Individually(b *testing.B) {
	for i := 0; i < b.N; i++ {
    for j := 100; j > 0; j-- {
  		FILE.Append(func(w *LineWriter) (e error) {
        w.WriteLine("1")
        return
      })
    }
	}
}

func BenchmarkAppendLine100(b *testing.B) {
	for i := 0; i < b.N; i++ {
    for j := 100; j > 0; j-- {
  		FILE.AppendLine("1")
    }
	}
}

func BenchmarkAppend1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FILE.Append(func(w *LineWriter) (e error) {
      for j := 1000; j > 0; j-- {
        w.WriteLine("1")
      }
      return
    })
	}
}

func BenchmarkAppend1000Individually(b *testing.B) {
	for i := 0; i < b.N; i++ {
    for j := 1000; j > 0; j-- {
  		FILE.Append(func(w *LineWriter) (e error) {
        w.WriteLine("1")
        return
      })
    }
	}
}

func BenchmarkAppendLine1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
    for j := 1000; j > 0; j-- {
  		FILE.AppendLine("1")
    }
	}
}
