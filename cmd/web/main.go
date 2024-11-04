package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", getHome)
	mux.HandleFunc("GET /snippet/view/{id}", GetSnippetView)
	mux.HandleFunc("GET /snippet/create", GetSnippetCreate)
	mux.HandleFunc("POST /snippet/create", PostSnippetCreate)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
