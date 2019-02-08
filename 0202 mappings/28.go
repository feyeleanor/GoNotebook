package main
import . "fmt"
import "sort"

type IntMap map[int] int
type Keys []int
type OrderedMap struct {
  Map map[int] int
  Keys
}

func main() {
  m1 := IntMap{0: 9, 1: 7, 2: 5, 3: 3, 4: 1}
  m2 := map[int] int{0: 0, 1: 2, 2: 4, 3: 6, 4: 8}
  m3 := NewOrderedMap(m1)
  m4 := NewOrderedMap(m2)

  for _, v := range []interface{} {m1, m2, m3, &m4} {
    print_elements(v)
  }
}

func print_elements(m interface{}) {
  switch m := m.(type) {
  case map[int] int:
    for i, v := range m {
      Printf("%v: %v\n", i, v)
    }
  case IntMap:
    print_elements(as_unnamed_map(m))
  case OrderedMap:
    for _, k := range m.Keys {
      Printf("%v: %v\n", k, m.Map[k])
    }
  case *OrderedMap:
    print_elements(*m)
  }
}

func as_unnamed_map(m map[int] int) map[int] int {
  return m
}

func NewOrderedMap(m map[int] int) (r OrderedMap) {
  r.Map = m
  for k := range r.Map {
    r.Keys = append(r.Keys, k)
  }
  sort.Sort(r.Keys)
  return
}

func (k Keys) Len() int {
  return len(k)
}

func (k Keys) Less(i, j int) bool {
  return k[i] > k[j]
}

func (k Keys) Swap(i, j int) {
  k[i], k[j] = k[j], k[i]
}
