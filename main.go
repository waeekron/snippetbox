package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Check if the path is an exact match, if not return 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello world!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {

	// Create a new servemux a.k.a router
	mux := http.NewServeMux()
	// Register a handler to the root router; maps / -> home -handler
	mux.HandleFunc("/", home)

	// Register snippet view and crate routes
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	// Start the web server, pass in the TCP network addr and servemux
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)

	}
}
