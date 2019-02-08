package main
import . "fmt"

type SparseArray struct {
  m map[int] int
  Count int
}

func main() {
  s := NewSparseArray(map[int] int{2: 4, 4: 8, 6: 12, 8: 16})
  print_sparse_array(s)
  s.Count = s.Count * 2
  print_sparse_array(s)
}

func print_sparse_array(s *SparseArray) {
  s.Range(func(i, v int) {
    Printf("%v: %v\n", i, v)
  })
}

func NewSparseArray(m map[int] int) (r *SparseArray) {
  r = &SparseArray{m: m}
  for k, _ := range m {
    if k > r.Count {
      r.Count = k
    }
  }
  return
}

func (s SparseArray) Range(f func(int, int)) {
  for i := 0; s.Count > -1; i++ {
    f(i, s.m[i])
    s.Count--
  }
}
