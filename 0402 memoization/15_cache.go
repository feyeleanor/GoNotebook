package main

import "math/big"

type Cache interface {
  Prefill(...*big.Int) Cache
  Load() Cache
  Fetch(*big.Int) (*big.Int, error)
  Store(*big.Int, *big.Int) (error)
  UseFor(f func(Cache) int) int
}
