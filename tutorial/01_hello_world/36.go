package main
import (
  . "fmt"
  . "net/http"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
  message := "hello world"
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, message)
  })

  done := make(chan bool)
  go func() {
    ListenAndServe(ADDRESS, nil)
    done <- true
  }()

  go func () {
    ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
    done <- true
  }()
  <- done
  <- done
}