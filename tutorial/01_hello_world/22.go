package main
import (
  . "fmt"
  . "net/http"
)

const MESSAGE = "hello world"
const ADDRESS = ":1024"

func main() {
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, MESSAGE)
  })
  ListenAndServe(ADDRESS, nil)
}