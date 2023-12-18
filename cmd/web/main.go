package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/waeekron/snippetbox/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {

	// read cli flags for port number and mysql dsn string
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:Miisukissa1!@/snippetbox?parseTime=true", "MySQL data sourcename")

	// parse cli flags
	flag.Parse()

	// infologger which logs to standard out
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// errorLoger to log errors to standard error out
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}

	// Close the connection pool before main() exists
	defer db.Close()

	// initialize a new application struct instance, containing the deps
	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	// start the server on localhost:4000/addr
	infoLog.Printf("Starting the server on %s", *addr)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}

// wraps sql.Open() and returns *sql.DB connection pool for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
