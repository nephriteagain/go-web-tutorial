//go:build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("HTTP Request: %+v\n", r)
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
	fmt.Println("server is running...")
    log.Fatal(http.ListenAndServe(":8080", nil))	
}