package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Access")
	flag.Parse()

	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", homeGet)
	router.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	router.HandleFunc("GET /snippet/create", snippetCreateGet)
	router.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, router)
	log.Fatal(err)
}
