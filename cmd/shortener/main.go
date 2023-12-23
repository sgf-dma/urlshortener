package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var shortenedUrls = map[string]string{}

func generateShortenedUrl() string {
	return "http://localhost:8080/" +
		GenerateString(len("EwHXdJfB"), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func checkIfItsUrl(s string) bool {
	_, err := url.Parse(s)
	if err != nil {
		return false
	}
	return true
}

func RootPageHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "Only POST requests are allowed to /", http.StatusBadRequest)
		return
	}

	if "text/plain" != req.Header.Get("Content-Type") {
		http.Error(res, "Content type must be text/plain", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		http.Error(res, "Failed to read the request body", http.StatusInternalServerError)
	}
	bodyString := string(body)
	if !checkIfItsUrl(bodyString) {
		http.Error(res, "Incorrect url format", http.StatusBadRequest)
	}
	shortenedUrl, alreadyShortenedUrl := shortenedUrls[bodyString]
	if !alreadyShortenedUrl {
		shortenedUrl := generateShortenedUrl()
		shortenedUrls[bodyString] = shortenedUrl
		res.WriteHeader(http.StatusCreated)
	}

	shortenedUrl, _ = shortenedUrls[bodyString]
	res.Header().Add("Content-Type", "text/plain")
	res.Header().Add("Content-Length", fmt.Sprintf("%d", len(shortenedUrl)))
	res.Write([]byte(shortenedUrl))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", RootPageHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
