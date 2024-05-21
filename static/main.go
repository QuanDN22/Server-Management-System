package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World! localhost:3000")
	port := flag.String("p", "3000", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
