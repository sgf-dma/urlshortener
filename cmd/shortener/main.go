package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func CheckIfItsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func RootPageHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed to /", http.StatusBadRequest)
		return
	}

	if req.Header.Get("Content-Type") != "text/plain" {
		http.Error(res, "Content type must be text/plain", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Failed to read the request body", http.StatusInternalServerError)
		return

	}
	bodyString := string(body)
	if !CheckIfItsURL(bodyString) {
		http.Error(res, "Incorrect url format", http.StatusBadRequest)
		return
	}

	shortenedURL := GetShortenedURL(bodyString)

	res.WriteHeader(http.StatusCreated)
	res.Header().Add("Content-Type", "text/plain")
	res.Header().Add("Content-Length", fmt.Sprintf("%d", len(shortenedURL)))
	res.Write([]byte(shortenedURL))
}

func IDHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Only GET requests are allowed to /{id}", http.StatusBadRequest)
		return
	}

	if req.Header.Get("Content-Type") != "text/plain" {
		http.Error(res, "Content type must be text/plain", http.StatusBadRequest)
		return
	}

	uri := req.RequestURI
	if MatchesGeneratedURLFormat(uri) {
		fullURL := GetFullURL(uri)
		if len(fullURL) > 0 {
			res.WriteHeader(http.StatusTemporaryRedirect)
			res.Write([]byte(fullURL))
		} else {
			http.Error(res, "Url not found", http.StatusNotFound)
		}
	} else {
		http.Error(res, "Invalid url format", http.StatusBadRequest)
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/shorten/", RootPageHandler)
	mux.HandleFunc("/", IDHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
