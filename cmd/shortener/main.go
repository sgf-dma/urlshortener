package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var shortenedUrls = map[string]string{}

func generateShortenedURL() string {
	return "http://localhost:8080/" +
		GenerateString(len("EwHXdJfB"), "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func checkIfItsURL(s string) bool {
	_, err := url.Parse(s)
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
	}
	bodyString := string(body)
	if !checkIfItsURL(bodyString) {
		http.Error(res, "Incorrect url format", http.StatusBadRequest)
	}
	_, alreadyShortenedURL := shortenedUrls[bodyString]
	if !alreadyShortenedURL {
		shortenedURL := generateShortenedURL()
		shortenedUrls[bodyString] = shortenedURL
		res.WriteHeader(http.StatusCreated)
	}

	shortenedURL := shortenedUrls[bodyString]
	res.Header().Add("Content-Type", "text/plain")
	res.Header().Add("Content-Length", fmt.Sprintf("%d", len(shortenedURL)))
	res.Write([]byte(shortenedURL))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", RootPageHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
