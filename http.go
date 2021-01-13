package main

import (
    "fmt"
    "net/http"
)

func HttpServe() {
    http.Handle("/", http.FileServer(http.Dir(*fe)))
    go http.ListenAndServe(fmt.Sprintf(":%d",*port), nil)
    fmt.Printf("HTTP server port: %d\n",*port)
    fmt.Printf("Frontend static files root path: %s\n",*fe)
}
