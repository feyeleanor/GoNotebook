package main
import (
  . "fmt"
  "net/http"
)

const MESSAGE = "hello world"
const ADDRESS = ":1024"

func main() {
  http.HandleFunc("/hello", Hello)
  http.ListenAndServe(ADDRESS, nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  Fprintf(w, MESSAGE)
}