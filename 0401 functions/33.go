package main

var limit int = 10

func main() {
  limit--
  if limit > 0 {
    main()
  }
}
