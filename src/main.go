package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HolaServidor)
    http.ListenAndServe(":8080", nil)
}

func HolaServidor(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Memo Docker")
}
