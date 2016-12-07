package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	dir             string
	prefix          string
	port            int
	allowDirListing bool
)

func init() {
	flag.StringVar(&dir, "dir", ".", "the directory to serve (/webroot)")
	flag.StringVar(&prefix, "prefix", "", "strip this prefix")
	flag.IntVar(&port, "port", 80, "the port to serve on")
	flag.BoolVar(&allowDirListing, "allow-directory-listing", false, "whether to allow directory listing")
	flag.Parse()
}

type justFilesFilesystem struct {
	Fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.Fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, _ := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}

func main() {
	spec := fmt.Sprintf("0.0.0.0:%d", port)
	logger := log.New(os.Stdout, "[serve] ", log.LstdFlags)
	logger.Printf("Serving %s on %s", dir, spec)
	var root http.FileSystem
	root = http.Dir(dir)
	if !allowDirListing {
		root = justFilesFilesystem{root}
	}
	handler := http.FileServer(root)
	if prefix != "" {
		handler = http.StripPrefix(prefix, handler)
	}
	handler = LoggerHandler{handler, logger}
	err := http.ListenAndServe(spec, handler)
	if err != nil {
		log.Fatalf("failed serving: %s", err)
	}
}
