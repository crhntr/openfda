package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"

	"github.com/crhntr/openfda/drug"
	"github.com/globalsign/mgo"
)

func importArgsInFiles() {
	session, err := mgo.Dial(":27017")
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, fileName := range flag.Args()[2:] {
		if err := func(i int, filename string) error {
			log.Printf("%d %q", i, fileName)

			f, err := os.Open(fileName)
			if err != nil {
				return err
			}
			defer f.Close()

			dec := json.NewDecoder(f)

			for {
				t, err := dec.Token()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return err
				}
				// fmt.Printf("%T %v\n", t, t)
				if str, ok := t.(string); ok && str == "results" {
					t, err = dec.Token()
					if err != nil {
						return err
					}
					if val, ok := t.(json.Delim); ok && val == json.Delim('[') {
						break
					}
				}
			}

			for {
				var rawEvent drug.RawEvent
				if err := dec.Decode(&rawEvent); err != nil {
					return err
				}

				event := rawEvent.Event()
				event.FileName = fileName

				if err := session.DB("openfda").C("drug_event").Insert(event); err != nil {
					if !mgo.IsDup(err) {
						log.Printf("%d %s %s %s", i, event.SafetyReportID, filename, err)
					}
				}

				if !dec.More() {
					return nil
				}
			}
			return nil
		}(i, fileName); err != nil {
			log.Printf("%d %s %s", i, fileName, err)
		}
	}
}

func importFiles() {
	session, err := mgo.Dial(":27017")
	if err != nil {
		log.Fatal(err)
		return
	}

	for i, fname := range jsonFiles {
		if err := func(i int, filename string) error {
			log.Printf("%d %q", i, filename)

			f, err := os.Open(filename)
			if err != nil {
				return err
			}
			defer f.Close()

			dec := json.NewDecoder(f)

			for {
				t, err := dec.Token()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return err
				}
				// fmt.Printf("%T %v\n", t, t)
				if str, ok := t.(string); ok && str == "results" {
					t, err = dec.Token()
					if err != nil {
						return err
					}
					if val, ok := t.(json.Delim); ok && val == json.Delim('[') {
						break
					}
				}
			}

			for {
				var rawEvent drug.RawEvent
				if err := dec.Decode(&rawEvent); err != nil {
					return err
				}

				event := rawEvent.Event()
				event.FileName = filename

				if err := session.DB("openfda").C("drug_event").Insert(event); err != nil {
					if !mgo.IsDup(err) {
						log.Printf("%d %s %s %s", i, event.SafetyReportID, filename, err)
					}
				}

				if !dec.More() {
					return nil
				}
			}
			return nil
		}(i, fname); err != nil {
			log.Printf("%d %s %s", i, fname, err)
		}
	}
}

func importAll() {
	session, err := mgo.Dial(":27017")
	if err != nil {
		log.Fatal(err)
		return
	}

	eachFile(func(f *os.File) error {
		fileName := f.Name()
		log.Printf("-> %q", fileName)

		dec := json.NewDecoder(f)

		for {
			t, err := dec.Token()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			// fmt.Printf("%T %v\n", t, t)
			if str, ok := t.(string); ok && str == "results" {
				t, err = dec.Token()
				if err != nil {
					return err
				}
				if val, ok := t.(json.Delim); ok && val == json.Delim('[') {
					break
				}
			}
		}

		for {
			var rawEvent drug.RawEvent
			if err := dec.Decode(&rawEvent); err != nil {
				return err
			}

			event := rawEvent.Event()
			event.FileName = fileName

			if err := session.DB("openfda").C("drug_event").Insert(event); err != nil {
				if !mgo.IsDup(err) {
					log.Printf("  %s %s %s", event.SafetyReportID, fileName, err)
				}
			}

			if !dec.More() {
				return nil
			}
		}
		return nil
	})
}
