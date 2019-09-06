package main

import . "os"

const FILE = "15.txt"

func main() {
  Exit(
    CachedFactorials(
      NewDiskCache(FILE),
      CalculateFactorials(Args[1:]...)))
}
