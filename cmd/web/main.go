package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Access")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	router := http.NewServeMux()
	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static/")})

	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", homeGet)
	router.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	router.HandleFunc("GET /snippet/create", snippetCreateGet)
	router.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("Starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, router)

	logger.Error(err.Error())
	os.Exit(1)
}
