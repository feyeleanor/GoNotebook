package main

import "bufio"
import "fmt"
import "os"
import "sync"
import "time"

const FILE LineFile = "21.txt"

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

func WriteToFile(n string, i int, f func()) {
  FILE.Append(func(w *LineWriter) (e error) {
    fmt.Printf("\t%v: writing at %v\n", n, time.Now().Second())
    _, e = w.WriteLine(fmt.Sprintf("%v: %v", n, i))
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

func (t *TaskScheduler) atomicWrite(n string, i int, f func()) {
  fmt.Printf("%v: acquiring lock\n", n)
  t.Lock()
  defer t.Unlock()
  defer fmt.Printf("%v: lock released\n", n)
  fmt.Printf("%v: lock acquired\n", n)
  WriteToFile(n, i, f)
}

func (t *TaskScheduler) CueTask(n string, l int, f func()) {
  t.tasks = append(t.tasks, func() {
    var w func(int)
    w = func(i int) {
      if i > 0 {
        t.atomicWrite(n, i, f)
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
  var t TaskScheduler
  t.CueTask("A", 6, func() {
    fmt.Println("\tA: sleeping for 2 seconds")
    time.Sleep(2 * time.Second)
  })
  t.CueTask("B", 6, func() {
    fmt.Println("\tB: sleeping for 1 second")
    time.Sleep(1 * time.Second)
  })
  t.Run()
}
