package main

import . "os"
import . "strconv"

var limit int

func init() {
  if len(Args) > 1 {
    if x, e := Atoi(Args[1]); e == nil {
      limit = x
    }
  }
}

func main() {
  limit--
  if limit > 0 {
    main()
  }
}
