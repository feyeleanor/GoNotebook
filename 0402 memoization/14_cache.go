package main

import "math/big"
import "bufio"
import "io"
import "os"

type Cache struct {
  memory map[string] string
  name string
  *bufio.Writer
}

func NewCache(s string) (r *Cache) {
  return &Cache{
    memory: make(map[string] string),
    name: s,
  }
}

func (c *Cache) Prefill(v ...*big.Int) *Cache {
  switch len(v) {
  case 0:
    return c
  case 1:
    c.memory[encode(v[0])] = "0"
    return c
  default:
    c.memory[encode(v[0])] = encode(v[1])
  }
  return c.Prefill(v[2:]...)
}

func (c *Cache) loadLine(f *bufio.Reader, n *big.Int) (e error) {
  var x string
  switch x, e = f.ReadString('\n'); {
  case len(x) == 0:
  case e == nil, e == io.EOF:
    c.memory[encode(n)] = x[:len(x) - 1]
  }
  return
}

func (c *Cache) Load() *Cache {
  f, _ := os.Open(c.name)
  defer f.Close()

  r := bufio.NewReader(f)
  Range(big.NewInt(1), func(n *big.Int) (e error) {
    return c.loadLine(r, n)
  })
  return c
}

func (c *Cache) Fetch(n *big.Int) (r *big.Int, e error) {
  if q, ok := c.memory[encode(n)]; ok {
    r, e = decode(q)
  }
  return
}

func (c *Cache) Store(n, v *big.Int) (e error) {
  x := encode(v)
  if _, e = c.WriteString(x + "\n"); e == nil {
    c.memory[encode(n)] = x
  }
  return
}

func (c *Cache) UseFor(f func(*Cache) int) int {
  o, _ := os.OpenFile(c.name, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0755)
  defer func() {
    c.Flush()
    o.Close()
  }()

  c.Writer = bufio.NewWriter(o)
  return f(c)
}
