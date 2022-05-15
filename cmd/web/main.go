package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {

	// address flag. ("flagName", "default value", "name for identification")
	// the flag import allows you take command line arguments from the OS.
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// must be called after all flags are defined and before they are accessed by the program.
	flag.Parse()

	// create a new logger for writing information messages and error messages.
	// parameters include (destination to write to,prefix for messages,flags to indicate additional info to include)
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/snippet", showSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)

	// file server that serves content from ui/static directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	/*
	Initialize a new http.Server struct and set Addr and handler fields 
	so server uses the same network address and routes as before, and set
	the ErrorLog field so that the server now uses the custom errorLog logger in
	the event of any problems.
	*/

	srv := &http.Server{
		Addr:     *addr,
		Handler:  mux,
		ErrorLog: errorLog,
	}

	infoLog.Printf("Starting server on port %s", *addr)

	err := srv.ListenAndServe()

	errorLog.Fatal(err)

}
