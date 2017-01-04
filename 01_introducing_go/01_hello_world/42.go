package main
import (
  . "fmt"
  . "net/http"
  "sync"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

var servers sync.WaitGroup

func main() {
  message := "hello world"
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, message)
  })

  Launch(func() {
    ListenAndServe(ADDRESS, nil)
  })

  Launch(func() {
    ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
  })
  servers.Wait()
}

func Launch(f func()) {
  servers.Add(1)
  go func() {
    defer servers.Done()
    f()
  }()
}