package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func GetSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display specific snippet with id %d...", id)
	w.Write([]byte(msg))
}

func GetSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func PostSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet..."))
}

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
