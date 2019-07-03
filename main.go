package main

import(
  "os"
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello world")
  })
  http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, os.Getenv("MY_LABEL_VERSION"))
  })
  http.HandleFunc("/nodename", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, os.Getenv("MY_NODE_NAME"))
  })
  http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, os.Getenv("MY_POD_NAME"))
  })
  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ok")
  })
  
  http.ListenAndServe(":8080", nil)
}

