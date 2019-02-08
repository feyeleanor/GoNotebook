package main

func main() {
  defer func() {
    recover()
  }()
  main()
}
