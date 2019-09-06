package main

import . "os"

func main() {
  Exit(
    CachedFactorials(
      MemoryCache{ "0": "1" },
      CalculateFactorials(Args[1:]...)))
}
