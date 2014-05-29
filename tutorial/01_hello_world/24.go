package main
import (
  . "fmt"
  . "net/http"
)

const ADDRESS = ":1024"

func main() {
  message := "hello world"
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, message)
  })
  ListenAndServeTLS(ADDRESS, "cert.pem", "key.pem", nil)
}