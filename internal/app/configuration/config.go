package configuration

import "flag"

type Options struct {
	ServerPort       uint
	ShortenerBaseUrl string
}

func ReadFlags() Options {
	var serverPort uint
	flag.UintVar(&serverPort, "a", 8080, "port on which the server should run")
	var shortenerBaseUrl string
	flag.StringVar(&shortenerBaseUrl, "b", "http://localhost:8080/", "specify base url")
	flag.Parse()
	return Options{
		ServerPort:       serverPort,
		ShortenerBaseUrl: shortenerBaseUrl,
	}
}
