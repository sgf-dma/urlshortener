package handlers

import (
	"errors"
	"github.com/Vla8islav/urlshortener/internal/app"
	"github.com/Vla8islav/urlshortener/internal/app/storage"
	"net/http"
)

func ExpandHandler(s storage.MakeshiftStorage, res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Only GET requests are allowed to /{id}", http.StatusBadRequest)
		return
	}

	uri := req.RequestURI
	if app.MatchesGeneratedURLFormat(uri) {
		fullURL, err := app.GetFullURL(s, uri)

		if err == nil {
			res.Header().Add("Location", fullURL)
			res.WriteHeader(http.StatusTemporaryRedirect)
		} else if errors.Is(err, app.ErrURLNotFound) {
			http.Error(res, "URL not found", http.StatusNotFound)
		} else {
			http.Error(res, "problem occured while extracting URL: "+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(res, "Invalid url format", http.StatusBadRequest)
	}
}
