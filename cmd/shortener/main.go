package main

import (
	"net/http"
	"net/url"
)

func CheckIfItsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", RootPageHandler)
	mux.HandleFunc("/geturl/", IDHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
