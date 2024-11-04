package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", homeGet)
	router.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	router.HandleFunc("GET /snippet/create", snippetCreateGet)
	router.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}
