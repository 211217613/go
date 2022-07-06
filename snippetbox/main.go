package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define a new command-line flag with the name 'addr'.
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Initialize a server struct and overwrite the ErrorLog field. This way the server will use the new custom error logger instead of the regular log.Error
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("starting server on port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
