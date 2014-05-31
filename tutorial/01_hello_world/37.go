package main
import (
  . "fmt"
  . "net/http"
  "os"
  "os/signal"
  . "sync"
)

const ADDRESS = ":1024"
const SECURE_ADDRESS = ":1025"

var servers WaitGroup

func init() {
  go SignalHandler(make(chan os.Signal, 1))
}

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

func SignalHandler(c chan os.Signal) {
  signal.Notify(c, os.Interrupt, os.Kill)
  for {
    switch s := <- c; s {
    case os.Interrupt:
      Println("^C received")
    case os.Kill:
      Println("SIGKIL received")
      os.Exit(1)
    }
  }
}