// Serve - quickly serve up a directory over HTTP
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8080", "port to listen on")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		usage()
		os.Exit(1)
	}
	http.Handle("/", http.FileServer(http.Dir(os.Args[1])))

	log.Fatalf("%s\n", http.ListenAndServe(net.JoinHostPort("", port), nil))
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "serve [-p 8080] directory\n\n")
	flag.PrintDefaults()
}
