package main

import (
    "fmt"
    "net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Pong from service 1!")
}

func main() {
    http.HandleFunc("/ping", pingHandler)
    http.ListenAndServe(":8001", nil)
}
