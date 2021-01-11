package main

import (
    "fmt"
    "net/http"
)

func HttpServe() {
    http.Handle("/", http.FileServer(http.Dir("fe/build/web")))
    go http.ListenAndServe(fmt.Sprintf(":%d",*port), nil)
    fmt.Println(fmt.Sprintf("HTTP server port: %d",*port))
}
