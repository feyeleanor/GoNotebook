package main
import . "fmt"
import . "net/http"
import "os"
import "sync"

var (
  address string
  secure_address string
  certificate string
  key string
)
var servers sync.WaitGroup

func init() {
  if address = os.Getenv("SERVE_HTTP"); address == "" {
    address = ":1024"
  }

  if secure_address = os.Getenv("SERVE_HTTPS"); secure_address == "" {
    secure_address = ":1025"
  }

  if certificate = os.Getenv("SERVE_CERT"); certificate == "" {
    certificate = "cert.pem"
  }

  if key = os.Getenv("SERVE_KEY"); key == "" {
    key = "key.pem"
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
    ListenAndServeTLS(secure_address, certificate, key, nil)
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
