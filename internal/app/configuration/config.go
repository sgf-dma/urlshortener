package configuration

import "flag"

type Options struct {
	ServerAddress    string
	ShortenerBaseURL string
}

var instance *Options

func ReadFlags() *Options {
	if instance == nil {
		opt := Options{}
		flag.StringVar(&opt.ServerAddress, "a", "localhost:8889", "port on which the server should run")
		flag.StringVar(&opt.ShortenerBaseURL, "b", "http://localhost:8000", "base url for shortened links")
		flag.Parse()
		instance = &opt
	}
	return instance
}
