package main

import (
	"net/url"
	"regexp"
	"strings"
)

var urlToShort = map[string]string{}
var shortToURL = map[string]string{}

const AllowedSymbolsInShortnedURL = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const GeneratedShortenedURLSample = "EwHXdJfB"

func GetShortenedURL(urlToShorten string) string {
	shortenedURL, err := GenerateShortenedURL()
	if err != nil {
		return ""
	}
	urlToShort[urlToShorten] = shortenedURL
	shortToURL[shortenedURL] = urlToShorten

	return shortenedURL
}

func GetFullURL(shortenedPostfix string) string {
	fullPath, err := url.JoinPath("http://localhost:8080/", shortenedPostfix)
	if err != nil {
		return ""
	}

	return shortToURL[fullPath]
}

func GenerateShortenedURL() (string, error) {
	fullPath, err := url.JoinPath("http://localhost:8080/",
		GenerateString(len(GeneratedShortenedURLSample), AllowedSymbolsInShortnedURL))
	if err != nil {
		return fullPath, err
	}
	return fullPath, nil
}

func MatchesGeneratedURLFormat(s string) bool {
	s = strings.Trim(s, "/")
	r, _ := regexp.Compile("^[" + AllowedSymbolsInShortnedURL + "]+$")
	return len(s) == len(GeneratedShortenedURLSample) && r.MatchString(s)
}
