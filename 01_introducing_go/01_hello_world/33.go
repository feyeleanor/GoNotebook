package main
import (
  . "fmt"
  . "net/http"
)

const SECURE_ADDRESS = ":1025"

func main() {
  message := "hello world"
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, message)
  })
  ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
}