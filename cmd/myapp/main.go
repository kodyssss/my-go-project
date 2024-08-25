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
    http.HandleFunc("/", helloHandler)
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
