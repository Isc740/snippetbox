package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", app.getHome)
	router.HandleFunc("GET /snippet/view/{id}", app.getSnippetView)
	router.HandleFunc("GET /snippet/create", app.getSnippetCreate)
	router.HandleFunc("POST /snippet/create", app.PostSnippetCreate)

	return router
}
