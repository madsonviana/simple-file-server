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
	log.Printf("Requisição: Método %s, URL: %v\n", r.Method, r.URL)
	c.fileServer.ServeHTTP(w, r)
}

var port = flag.String("port", "9999", "A porta http")

func main() {
	flag.Parse()

	log.Printf("Iniciando file server na porta %s\n", *port)

	log.Fatal(http.ListenAndServe(
		":"+*port,
		&crossHandler{http.FileServer(http.Dir("./"))},
	))
}
