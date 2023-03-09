package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// read cli flag which defines the port number
	addr := flag.String("addr", ":4000", "HTTP network address")

	// parse cli flags
	flag.Parse()

	// infologger which logs to standard out
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// errorLoger to log errors to standard error out
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize a new application struct instance, containing the deps
	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}
	// start the server on localhost:4000/addr
	infoLog.Printf("Starting the server on %s", *addr)
	// call the ListenAndServe() on our http.Server struct:
	err := srv.ListenAndServe()
	errLog.Fatal(err)
}
