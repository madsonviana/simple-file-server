package main

import (
	"flag"
	"log"
	"net/http"
)

type crossHandler struct {
	fileServer http.Handler
}

func (c *crossHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Printf("New Request: Method %s, URL: %v\n", r.Method, r.URL)
	c.fileServer.ServeHTTP(w, r)
}

var port = flag.String("port", "9999", "The http port to listen")

func main() {
	flag.Parse()

	log.Printf("Starting server on port %s\n", *port)

	log.Fatal(http.ListenAndServe(
		":"+*port,
		&crossHandler{http.FileServer(http.Dir("./"))},
	))
}
