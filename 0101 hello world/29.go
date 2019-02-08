package main
import . "fmt"
import "net/http"

const MESSAGE = "hello world"
const ADDRESS = ":1024"

func main() {
  http.HandleFunc("/hello", Hello)
  if e := http.ListenAndServe(ADDRESS, nil); e != nil {
    Println(e)
  }
}

func Hello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain")
  Fprintf(w, MESSAGE)
}
