package main

import (
	"github.com/Vla8islav/urlshortener/internal/app/configuration"
	"github.com/Vla8islav/urlshortener/internal/app/handlers"
	"github.com/Vla8islav/urlshortener/internal/app/storage"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	s := storage.NewInstance()
	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
	    handlers.RootPageHandler(s, w, r)
	})
	r.HandleFunc("/{slug:[A-Za-z]+}", func(w http.ResponseWriter, r *http.Request) {
	    handlers.ExpandHandler(s, w, r)
	})

	err := http.ListenAndServe(configuration.ReadFlags().ServerAddress, r)
	if err != nil {
		panic(err)
	}
}
