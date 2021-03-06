package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/crhntr/openfda/drug"
)

func createTestData(search string) {
	var events []drug.Event

	ensureDataDir()

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

	type LabelRequestResponse struct {
		Meta struct {
			Results struct {
				Limit int `json:"limit"`
				Skip  int `json:"skip"`
				Total int `json:"total"`
			} `json:"results"`
		} `json:"meta"`

		Results []drug.Label `json:"results"`
	}

	skip := 0

	for {
		reqURL := joinURL(apiURL, "drug", "event.json") + "?limit=100&search=" + search
		if skip > 0 {
			reqURL += fmt.Sprintf("&skip=%d", skip)
		}

		// log.Printf("requesting %q", reqURL)
		req, err := http.Get(reqURL)
		if err != nil {
			log.Fatalf("drug event request could not be made: %q", err)
		}
		if req.StatusCode != http.StatusOK {
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

	// var drugs drug.EventDrug // must be deduped

	// map key should be drug's SPLSetID
	labels := make(map[string]drug.Label)

	for _, event := range events {
		for _, drg := range event.Drugs {
			if drg.OpenFDA.IsHumanPresecriptionDrug() &&
				len(drg.OpenFDA.BrandName) == 1 &&
				len(drg.OpenFDA.SPLSetID) > 0 {

				labels[drg.OpenFDA.SPLSetID[0]] = drug.Label{}
			}
		}
	}

	for splSetID, _ := range labels {
		lreqURL := joinURL(apiURL, "drug", "label.json") + "?search=openfda.spl_set_id:%22" + splSetID + "%22"
		lreq, err := http.Get(lreqURL)
		if err != nil {
			log.Printf("drug event request could not be made: %q", err)
			continue
		}
		if lreq.StatusCode == http.StatusNotFound {
			continue
		}
		if lreq.StatusCode != http.StatusOK {
			log.Printf("drug event request not successful %q. Got http status %d\n\n", lreqURL, lreq.StatusCode)
			io.Copy(os.Stderr, lreq.Body)
			continue
		}

		var labelRes LabelRequestResponse

		if err := json.NewDecoder(lreq.Body).Decode(&labelRes); err != nil {
			log.Printf("could not parse json response for %q: %q", lreq.Request.URL, err)
			continue
		}

		if len(labelRes.Results) > 0 {
			labels[splSetID] = labelRes.Results[0]
		}
	}

	fmt.Println(len(labels))
	fmt.Println(len(events))

	for key, label := range labels {
		if key == "" {
			continue
		}
		outp := path.Join(*outPath, "drug", "label", key+".json")

		f, err := os.Create(outp)
		if err != nil {
			if !os.IsExist(err) {
				f, err = os.Open(outp)
			}
			if err != nil {
				log.Printf("Error creating label file for %q: %q", key, err)
			}
		}

		if err := json.NewEncoder(f).Encode(label); err != nil {
			log.Printf("Error encoding json for label %q: %q", key, err)
		}
		f.Close()
	}

	for _, event := range events {
		if event.SafetyReportID == "" {
			continue
		}

		outp := path.Join(*outPath, "drug", "event", event.SafetyReportID+".json")

		f, err := os.Create(outp)
		if err != nil {
			if !os.IsExist(err) {
				f, err = os.Open(outp)
			}
			if err != nil {
				log.Printf("Error creating event file for %q: %q", event.SafetyReportID, err)
			}
		}

		if err := json.NewEncoder(f).Encode(event); err != nil {
			log.Printf("Error encoding json for label %q: %q", event.SafetyReportID, err)
		}
		f.Close()
	}

	// insert labels and events into database

}
