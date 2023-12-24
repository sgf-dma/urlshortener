package main

import (
	"github.com/Vla8islav/urlshortener/internal/app/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RootPageHandler)
	mux.HandleFunc("/geturl/", handlers.IDHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
