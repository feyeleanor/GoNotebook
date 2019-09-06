package main

import "math/big"
import "bufio"
import "io"
import "os"

type DiskCache struct {
  memory map[string] string
  name string
  *bufio.Writer
}

func NewDiskCache(s string) (r *DiskCache) {
  return &DiskCache{
    memory: make(map[string] string),
    name: s,
  }
}

func (d *DiskCache) Prefill(v ...*big.Int) Cache {
  switch len(v) {
  case 0:
    return d
  case 1:
    d.memory[encode(v[0])] = "0"
    return d
  default:
    d.memory[encode(v[0])] = encode(v[1])
  }
  return d.Prefill(v[2:]...)
}

func (d *DiskCache) loadLine(f *bufio.Reader, n *big.Int) (e error) {
  var x string
  switch x, e = f.ReadString('\n'); {
  case len(x) == 0:
  case e == nil, e == io.EOF:
    d.memory[encode(n)] = x[:len(x) - 1]
  }
  return
}

func (d *DiskCache) Load() Cache {
  f, _ := os.Open(d.name)
  defer f.Close()

  r := bufio.NewReader(f)
  Range(big.NewInt(1), func(n *big.Int) (e error) {
    return d.loadLine(r, n)
  })
  return d
}

func (d *DiskCache) Fetch(n *big.Int) (r *big.Int, e error) {
  if q, ok := d.memory[encode(n)]; ok {
    r, e = decode(q)
  }
  return
}

func (d *DiskCache) Store(n, v *big.Int) (e error) {
  x := encode(v)
  if _, e = d.WriteString(x + "\n"); e == nil {
    d.memory[encode(n)] = x
  }
  return
}

func (d *DiskCache) UseFor(f func(Cache) int) int {
  o, _ := os.OpenFile(d.name, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0755)
  defer func() {
    d.Flush()
    o.Close()
  }()

  d.Writer = bufio.NewWriter(o)
  return f(d)
}
