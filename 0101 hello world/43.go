package main
import . "fmt"
import . "net/http"
import "os"
import "sync"

const SECURE_ADDRESS = ":1025"

var address string
var servers sync.WaitGroup

func init() {
  if address = os.Getenv("SERVE_HTTP"); address == "" {
    address = ":1024"
  }
}

func main() {
  message := "hello world"
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, message)
  })

  Launch(func() {
    ListenAndServe(address, nil)
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
