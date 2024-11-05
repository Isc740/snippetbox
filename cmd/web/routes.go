package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", app.homeGet)
	router.HandleFunc("GET /snippet/view/{id}", app.snippetViewGet)
	router.HandleFunc("GET /snippet/create", app.snippetCreateGet)
	router.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	return router
}
