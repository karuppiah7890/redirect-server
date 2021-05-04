package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage: redirect-server [listen-address] [redirection-location-url]")
	}

	listenAddress := os.Args[1]

	redirectionLocationUserInput := os.Args[2]
	redirectionLocation, urlParseErr := url.Parse(redirectionLocationUserInput)
	if urlParseErr != nil {
		log.Fatalf("Error parsing redirection location %v: %v", redirectionLocationUserInput, urlParseErr)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", redirectionLocation.String())
		w.WriteHeader(302)
	})

	listenError := http.ListenAndServe(listenAddress, nil)
	if listenError != nil {
		log.Fatal(urlParseErr)
	}
}
