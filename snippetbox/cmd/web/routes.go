package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Setup sql driver and connections
	mux := http.NewServeMux()

	// handle static assets
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
