package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    val := os.Getenv("APPLICATION_NAME")
    fmt.Fprintf(w, "Congratulations! Your Go application has been successfully deployed on Kubernetes. Your application name is %s", val)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}