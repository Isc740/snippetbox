package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Home", "Go")
	w.Write([]byte("Hello from snippetbox"))
}

func GetSnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display specific snippet with id %d...", id)
}

func GetSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func PostSnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create new snippet..."))
}
