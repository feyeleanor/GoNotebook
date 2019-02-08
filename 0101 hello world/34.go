package main
import . "fmt"
import . "net/http"

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

func main() {
  message := "hello world"
  HandleFunc("/hello", func(w ResponseWriter, r *Request) {
    w.Header().Set("Content-Type", "text/plain")
    Fprintf(w, message)
  })

  go func() {
    ListenAndServe(ADDRESS, nil)
  }()

  ListenAndServeTLS(SECURE_ADDRESS, "cert.pem", "key.pem", nil)
}
