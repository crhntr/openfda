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

func createTestData(seedEventID string) {
	var events []drug.Event

	type EventRequestResponse struct {
		Meta struct {
			Limit int `json:"limit"`
			Skip  int `json:"skip"`
			Total int `json:"total"`
		} `json:"meta"`

		Events []drug.RawEvent `json:"results"`
	}

	for {
		reqURL := joinURL(apiURL, "drug", "event.json") + "?search=safetyreportid.exact:%22" + seedEventID + "%22"
		seedEventReq, err := http.Get(reqURL)
		if err != nil {
			log.Fatalf("drug event request could not be made: %q", err)
		}
		if seedEventReq.StatusCode != 200 {
			fmt.Printf("drug event request not successful %q. Got http status %d\n\n", reqURL, seedEventReq.StatusCode)
			io.Copy(os.Stderr, seedEventReq.Body)
			os.Exit(1)
		}

		var data EventRequestResponse
		if err := json.NewDecoder(seedEventReq.Body).Decode(&data); err != nil {
			log.Fatalf("could not parse json response for %q: %q", data.Request.URL, err)
		}

		if data.Meta.Total <= data.Meta.Skip+data.Meta.Limit {
			break
		}

		for _, rEvent := range eventResponse.Events {
			events = append(events, rEvent.Event())
		}
	}

	for _, event := range events {
		for _, drg := range event.Drugs {
			if drg.OpenFDA.IsHumanPresecriptionDrug() {
				fmt.Println(drg.OpenFDA.BrandName)
			}
		}
	}

}
