package main

import "bufio"
import "fmt"
import "os"
import "sync"

const FILE LineFile = "22.txt"

type LineWriter struct {
  *bufio.Writer
}

func NewLineWriter(f *os.File) (r *LineWriter) {
  return &LineWriter{ bufio.NewWriter(f) }
}

func (l *LineWriter) WriteLine(s string) (int, error) {
  return l.WriteString(s + "\n")
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

func WriteToFile(i int, f func()) {
  FILE.Append(func(w *LineWriter) (e error) {
    _, e = w.WriteLine(fmt.Sprint(i))
    f()
    return
  })
}

func Launch(f []func()) {
  if len(f) > 0 {
    go f[0]()
    Launch(f[1:])
  }
}

type TaskScheduler struct {
  sync.Mutex
  sync.WaitGroup
  tasks []func()
}

func log(n, s string) {
  fmt.Printf("%v: %v\n", n, s)
}

func (t *TaskScheduler) Guard(n string, f func()) {
  log(n, "acquiring lock")
  t.Lock()
  defer log(n, "lock released")
  defer t.Unlock()
  defer log(n, "releasing lock")
  log(n, "lock acquired")
  f()
}

func (t *TaskScheduler) CueTask(n string, l int, f func(int)) {
  t.tasks = append(t.tasks, func() {
    var w func(int)
    w = func(i int) {
      if i > 0 {
        t.Guard(n, func() {
          f(i)
        })
        w(i - 1)
      }
    }
    w(l)
    t.Done()
  })
}

func (t *TaskScheduler) Run() {
  t.Add(len(t.tasks))
  Launch(t.tasks)
  t.Wait()
}

func main() {
  p := func(i int) {
    WriteToFile(i, func() {
      fmt.Println("\twrote", i)
    })
  }

  var t TaskScheduler
  t.CueTask("A", 4, p)
  t.CueTask("B", 4, p)
  t.CueTask("C", 4, p)
  t.Run()
}
