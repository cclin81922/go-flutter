package main

import (
    "github.com/cclin81922/go-3rd-party/go-flag"
)

var (
    port = flag.Int("port", 8080, "HTTP server port")
    fe = flag.String("fe", "fe/build/web", "Frontend static files root path")
)

func parseArgs() {
    flag.Parse()
}
