package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

var (
    dir  string
    port int
)

func init() {
    flag.StringVar(&dir, "dir", ".", "the directory to serve")
    flag.IntVar(&port, "port", 8080, "the port to serve on")
    flag.Parse()
}

func main() {
    spec := fmt.Sprintf("0.0.0.0:%d", port)
    log.Printf("Serving %s on %s", dir, spec)
    err := http.ListenAndServe(spec, http.FileServer(http.Dir(dir)))
    if err != nil {
        log.Fatalf("failed serving: %s", err)
    }
}
