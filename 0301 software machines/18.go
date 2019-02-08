package main
import "fmt"
import r "reflect"
import "unsafe"

type Memory []uintptr

var _BYTE_SLICE = r.TypeOf([]byte(nil))
var _MEMORY = r.TypeOf(Memory{})
var _MEMORY_BYTES = int(_MEMORY.Elem().Size())

func (m Memory) newHeader() (h r.SliceHeader) {
  h = *(*r.SliceHeader)(unsafe.Pointer(&m))
  h.Len = len(m) * _MEMORY_BYTES
  h.Cap = cap(m) * _MEMORY_BYTES
  return
}

func (m *Memory) Serialise() (b []byte) {
  h := m.newHeader()
  b = make([]byte, h.Len)
  copy(b, *(*[]byte)(unsafe.Pointer(&h)))
  return
}

func (m *Memory) Overwrite(i interface{}) {
  switch i := i.(type) {
  case Memory:
    copy(*m, i)
  case []byte:
    h := m.newHeader()
    b := *(*[]byte)(unsafe.Pointer(&h))
    copy(b, i)
  }
}
func (m *Memory) Bytes() (b []byte) {
  h := m.newHeader()
  return *(*[]byte)(unsafe.Pointer(&h))
}

func main() {
  m := make(Memory, 2)
  b := m.Bytes()
  s := m.Serialise()
  fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
  fmt.Println("b (bytes) =", len(b), "of", cap(b), ":", b)
  fmt.Println("s (bytes) =", len(s), "of", cap(s), ":", s)

  m.Overwrite(Memory{3, 5})
  fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
  fmt.Println("b (bytes) =", len(b), "of", cap(b), ":", b)
  fmt.Println("s (bytes) =", len(s), "of", cap(s), ":", s)

  s = m.Serialise()
  m.Overwrite([]byte{8, 7, 6, 5, 4, 3, 2, 1})
  fmt.Println("m (cells) =", len(m), "of", cap(m), ":", m)
  fmt.Println("b (bytes) =", len(b), "of", cap(b), ":", b)
  fmt.Println("s (bytes) =", len(s), "of", cap(s), ":", s)
}
