package main

import "bufio"
import "fmt"
import "os"
import "sync"

const FILE LineFile = "23.txt"

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

type TaskList []func()

func (t TaskList) Len() int {
  return len(t)
}

func (t TaskList) Launch() {
  if t.Len() > 0 {
    go t[0]()
    t[1:].Launch()
  }
}

type Lockable interface {
  Lock()
  Unlock()
}

type Scheduler struct {
  Lockable
  sync.WaitGroup
  TaskList
}

func NewScheduler(l Lockable) *Scheduler {
  return &Scheduler{ Lockable: l }
}

func log(n, s string) {
  fmt.Printf("%v: %v\n", n, s)
}

func (s *Scheduler) Guard(n string, f func()) {
  log(n, "acquiring lock")
  s.Lock()
  defer log(n, "lock released")
  defer s.Unlock()
  defer log(n, "releasing lock")
  log(n, "lock acquired")
  f()
}

func (s *Scheduler) CueTask(n string, l int, f func(int)) {
  s.TaskList = append(s.TaskList, func() {
    var w func(int)
    w = func(i int) {
      if i > 0 {
        s.Guard(n, func() {
          f(i)
        })
        w(i - 1)
      }
    }
    w(l)
    s.Done()
  })
}

func (s *Scheduler) Run() {
  s.Add(s.Len())
  s.Launch()
  s.Wait()
}

func main() {
  p := func(i int) {
    WriteToFile(i, func() {
      fmt.Println("\twrote", i)
    })
  }

  s := NewScheduler(new(sync.Mutex))
  s.CueTask("A", 4, p)
  s.CueTask("B", 4, p)
  s.CueTask("C", 4, p)
  s.Run()
}
