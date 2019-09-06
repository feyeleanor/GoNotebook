package main

import "math/big"
import . "fmt"
import . "os"
import . "strconv"

func main() {
  for _, s := range Args[1:] {
    if x, e := Atoi(s); e == nil {
      if b := MaxBigInt(x); e == nil {
        Printf("max value of %v words is %v\n", x, b)
      }
    }
  }
}

func MaxBigInt(n int) *big.Int {
  w := []big.Word{}
  for i := n; i > 0; i-- {
    w = append(w, 0xFFFFFFFFFFFFFFFF)
  }
  return new(big.Int).SetBits(w)
}
