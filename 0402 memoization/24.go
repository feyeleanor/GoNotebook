package main

import "bufio"
import "fmt"
import "os"
import "sync"

const FILE LineFile = "24.txt"

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
  s.Lock()
  defer s.Unlock()
  log(n, "lock acquired")
  f()
}

type Payload func(int) bool

func (s *Scheduler) CueTask(n string, p Payload) {
  s.TaskList = append(s.TaskList, func() {
    var w Payload
    w = func(i int) (ok bool) {
      s.Guard(n, func() {
        ok = p(i)
      })
      return ok && w(i + 1)
    }
    w(0)
    s.Done()
  })
}

func (s *Scheduler) Run() {
  s.Add(s.Len())
  s.Launch()
  s.Wait()
}

func main() {
  p := func(l int) Payload {
    return func(i int) (ok bool) {
      if ok = i < l; ok {
        i = l - i
        WriteToFile(i, func() {
          fmt.Println("\twrote", i)
        })
      }
      return
    }
  }

  s := NewScheduler(new(sync.Mutex))
  s.CueTask("A", p(4))
  s.CueTask("B", p(4))
  s.CueTask("C", p(4))
  s.Run()
}
