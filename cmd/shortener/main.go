package main

import (
	"github.com/Vla8islav/urlshortener/internal/app"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.RootPageHandler)
	mux.HandleFunc("/geturl/", app.IDHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
