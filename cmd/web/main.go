package main

import (
	"log"
	"net/http"
)

func main() {

	// create servemux
	mux := http.NewServeMux()

	// add supported URL's and their handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("snippet/create", snippetCreate)

	// start the server on localhost:4000
	log.Print("Server starting on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
