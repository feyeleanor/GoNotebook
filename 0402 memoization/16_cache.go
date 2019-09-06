package main

import "math/big"

type ReadCache interface {
  Fetch(*big.Int) (*big.Int, error)
}

type WriteCache interface {
  Store(*big.Int, *big.Int) (error)
}

type Cache interface {
  ReadCache
  WriteCache
}
