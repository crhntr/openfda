package main

import (
	"io"
	"log"
	"net/url"
)

const (
	apiURL = "https://api.fda.gov"

	downloadsFilePath = apiURL + "/download.json"
)

func apiRequest(endpoint string, search string, then func(r io.Reader) (bool, error)) {
	path := joinURL(apiURL, "drug", endpoint)
	reqURL, err := url.Parse(path)
	if err != nil {
		log.Fatalf("could not create url %q: %q", path, err)
	}
	reqURL.Query().Set("search", search)
	log.Print(reqURL)
}
