package main

import (
    "fmt"
    "log"
    "net/http"
    "my-go-project/pkg/mylib"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, mylib.Hello())
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}