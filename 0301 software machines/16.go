package main
import "fmt"

type AssocArray struct {
  Key   string
  Value interface{}
  Next  *AssocArray
}

func (a *AssocArray) GetIf(k string) (r interface{}) {
  if a != nil && a.Key == k {
    r = a.Value
  }
  return
}

type Search struct {
  Term         string
  Value        interface{}
  Cursor, Memo *AssocArray
}

func (s *Search) Step() *Search {
  s.Value = s.Cursor.GetIf(s.Term)
  return s
}

func (s *Search) Searching() bool {
  return s.Value == nil && s.Cursor != nil
}

func Find(a *AssocArray, k string) (s *Search) {
  s = &Search{Term: k, Cursor: a}
  for s.Step(); s.Searching(); s.Step() {
    s.Memo = s.Cursor
    s.Cursor = s.Cursor.Next
  }
  return
}

type Map []*AssocArray

func (m Map) Chain(k string) int {
  var c uint
  for i := len(k) - 1; i > 0; i-- {
    c = c << 8
    c += (uint)(k[i])
  }
  return int(c) % len(m)
}

func (m Map) Set(k string, v interface{}) {
  c := m.Chain(k)
  a := m[c]
  s := Find(a, k)
  if s.Value != nil {
    s.Cursor.Value = v
  } else {
    n := &AssocArray{Key: k, Value: v}
    switch {
    case s.Cursor == a:
      n.Next = s.Cursor
      m[c] = n
    case s.Cursor == nil:
      s.Memo.Next = n
    default:
      n.Next = s.Cursor
      s.Memo.Next = n
    }
  }
}

func (m Map) Get(k string) (r interface{}) {
  if s := Find(m[m.Chain(k)], k); s != nil {
    r = s.Value
  }
  return
}

func main() {
  m := make(Map, 1024)
  m.Set("apple", "rosy")
  fmt.Printf("%v\n", m.Get("apple"))

  m.Set("blueberry", "sweet")
  fmt.Printf("%v\n", m.Get("blueberry"))

  m.Set("cherry", "pie")
  fmt.Printf("%v\n", m.Get("cherry"))

  m.Set("cherry", "tart")
  fmt.Printf("%v\n", m.Get("cherry"))

  fmt.Printf("%v\n", m.Get("tart"))
}
