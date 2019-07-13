/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
*/
package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8888", "port to serve on")
	flag.Parse()

	directory := "."
	if flag.NArg() > 0 {
		directory = flag.Arg(0)
	}

	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving %s on HTTP port: %s\n", directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
