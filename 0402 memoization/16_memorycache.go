package main

import "math/big"

type MemoryCache map[string] string

func (m MemoryCache) Fetch(n *big.Int) (r *big.Int, e error) {
  if q, ok := m[encode(n)]; ok {
    r, e = decode(q)
  }
  return
}

func (m MemoryCache) Store(n, v *big.Int) (e error) {
  m[encode(n)] = encode(v)
  return
}
