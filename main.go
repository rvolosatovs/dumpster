package main

import (
	"flag"
	"log"
	"net/http"
)

var path = flag.String("path", "/var/lib/dumpster", "path to dump")
var addr = flag.String("http", "", "HTTP address to dump to")

func init() {
	log.SetFlags(0)
	flag.Parse()
}

func main() {
	if err := http.ListenAndServe(*addr, http.FileServer(http.Dir(*path))); err != nil {
		log.Fatalf("Failed to dump: `%s`", err)
	}
}
