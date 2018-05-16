package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/crhntr/openfda/drug"
)

func createTestData(search string) {
	var events []drug.Event

	type EventRequestResponse struct {
		Meta struct {
			Results struct {
				Limit int `json:"limit"`
				Skip  int `json:"skip"`
				Total int `json:"total"`
			} `json:"results"`
		} `json:"meta"`

		Events []drug.RawEvent `json:"results"`
	}

	skip := 0

	for {
		reqURL := joinURL(apiURL, "drug", "event.json") + "?limit=100&search=" + search
		if skip > 0 {
			reqURL += fmt.Sprintf("&skip=%d", skip)
		}

		log.Printf("requesting %q", reqURL)
		req, err := http.Get(reqURL)
		if err != nil {
			log.Fatalf("drug event request could not be made: %q", err)
		}
		if req.StatusCode != 200 {
			fmt.Printf("drug event request not successful %q. Got http status %d\n\n", reqURL, req.StatusCode)
			io.Copy(os.Stderr, req.Body)
			os.Exit(1)
		}

		var data EventRequestResponse
		if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
			log.Fatalf("could not parse json response for %q: %q", req.Request.URL, err)
		}

		skip += data.Meta.Results.Limit
		for _, rEvent := range data.Events {
			events = append(events, rEvent.Event())
		}

		if len(events) >= data.Meta.Results.Total {
			break
		}
	}

	for _, event := range events {
		for _, drg := range event.Drugs {
			if drg.OpenFDA.IsHumanPresecriptionDrug() && len(drg.OpenFDA.BrandName) == 1 {
				fmt.Println(drg.OpenFDA.BrandName)
			}
		}
	}

	// fetch relevant labels

	// insert labels and events into database

}
