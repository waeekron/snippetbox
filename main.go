package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	// Get the url query parameter id
	idStr := r.URL.Query().Get("id")
	// Try to convert the string value to int
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		//http.Error(w, "Not found", http.StatusNotFound)
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying snippet with id %d", id)

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
