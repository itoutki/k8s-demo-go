package main

import(
  "os"
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hello world")
  })
  http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, os.Getenv("MY_LABEL_VERSION"))
  })
  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "ok")
  })
  
  http.ListenAndServe(":8080", nil)
}

