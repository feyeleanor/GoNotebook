package main

import "math/big"
import . "fmt"

func main() {
  b := new(big.Int)
  b.SetBits([]big.Word{ 0xFFFFFFFFFFFFFFFF })
  Println("The largest number in 1 word is", b)
}
