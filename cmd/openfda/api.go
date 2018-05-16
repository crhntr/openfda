package main

import (
	"io"
	"net/http"
)

const (
	apiURL = "https://api.fda.gov"

	downloadsFilePath = apiURL + "/download.json"
)

func apiRequest(endpoint string, limit int, then func(r io.Reader)) {
	http.Get("")
}
