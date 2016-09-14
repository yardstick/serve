package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	dir    string
	prefix string
	port   int
)

func init() {
	flag.StringVar(&dir, "dir", ".", "the directory to serve (/webroot)")
	flag.StringVar(&prefix, "prefix", "", "strip this prefix")
	flag.IntVar(&port, "port", 80, "the port to serve on")
	flag.Parse()
}

func main() {
	spec := fmt.Sprintf("0.0.0.0:%d", port)
	logger := log.New(os.Stdout, "[serve] ", log.LstdFlags)
	logger.Printf("Serving %s on %s", dir, spec)
	handler := http.FileServer(http.Dir(dir))
	if prefix != "" {
		handler = http.StripPrefix(prefix, handler)
	}
	handler = LoggerHandler{handler, logger}
	err := http.ListenAndServe(spec, handler)
	if err != nil {
		log.Fatalf("failed serving: %s", err)
	}
}
