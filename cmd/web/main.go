package main

import (
	"database/sql"
	"flag"
	_"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// address flag. ("flagName", "default value", "name for identification")
	// the flag import allows you take command line arguments from the OS.
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "odin@/snippetbox?parseTime=true", "MySQL data source name")
	// must be called after all flags are defined and before they are accessed by the program.
	flag.Parse()
	// create a new logger for writing information messages and error messages.
	// parameters include (destination to write to,prefix for messages,flags to indicate additional info to include)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()

	// Initialize a new instance of the application containing dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	/*
		Initialize a new http.Server struct and set Addr and handler fields
		so server uses the same network address and routes as before, and set
		the ErrorLog field so that the server now uses the custom errorLog logger in
		the event of any problems.
	*/

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}
	infoLog.Printf("Starting server on port %s", *addr)
	errr := srv.ListenAndServe()
	errorLog.Fatal(errr)

}

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
