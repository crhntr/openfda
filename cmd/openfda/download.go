package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const dataDir = "openfda_data"

func ensureDataDir() {
	_, err := os.Stat(dataDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(dataDir, 0700)
		if err != nil {
			log.Fatal("could not create data directory %s: %q", dataDir, err)
		}
	}
}

func downloadsFile() {
	r, err := http.Get("https://api.fda.gov/download.json")
	if err != nil {
		log.Fatalf("could not get downloads file: %q", err)
	}

	ensureDataDir()

	downloadJSON, err := os.Create(path.Join(dataDir, "download.json"))
	if err != nil {
		log.Fatalf("could not create downloads file: %q", err)
	}

	io.Copy(downloadJSON, r.Body)
}

func download() {
	downloads := openDownloads()
	size := downloads.Results.Drug.Event.Size()
	var downlaoded float64

	parts := downloads.Results.Drug.Event.Partitions.Filter(*year, *quarter)
	fmt.Println(len(parts))
	for i, part := range parts {
		func(i int, part Partition) {
			dir := path.Dir(strings.TrimPrefix(part.File, "https://download.open.fda.gov/"))

			if err := os.MkdirAll(path.Join(*outPath, dir), os.ModePerm); err != nil {
				log.Printf("%d %s %s", i, part.File, err)
				return
			}

			log.Printf("  %d %q %.2f %.2f", i, part.Name, part.Size, (downlaoded/size)*100)
			resp, err := http.Get(part.File)
			if err != nil {
				log.Printf("  %d %s %s", i, part.File, err)
				return
			}
			defer resp.Body.Close()

			out, err := os.Create(path.Join(*outPath, dir, path.Base(part.File)))
			if err != nil {
				log.Printf("%d %s %s", i, part.File, err)
				return
			}
			defer out.Close()
			io.Copy(out, resp.Body)
			downlaoded += part.Size
		}(i, part)
	}
}

func openDownloads() Downloads {
	f, err := os.Open(path.Join(dataDir, "download.json")) // https://api.fda.gov/download.json
	if err != nil {
		panic(err)
	}
	var downloads Downloads
	if err = json.NewDecoder(f).Decode(&downloads); err != nil {
		panic(err)
	}
	return downloads
}

type Downloads struct {
	Results struct {
		Drug struct {
			Event DataSet `json:"event"`
			Label DataSet `json:"label"`
		} `json:"drug"`
	} `json:"results"`
}

type DataSet struct {
	TotalRecords int    `json:"total_records"`
	ExportDate   string `json:"export_date"`
	Partitions   `json:"partitions"`
}

type Partitions []Partition
type Partition struct {
	Size    float64 `json:"size_mb,string"`
	Records int     `json:"records"`
	Name    string  `json:"display_name"`
	File    string  `json:"file"`
}

func (parts Partitions) Filter(year, quarter string) Partitions {
	var filteredParts Partitions

	for _, part := range parts {
		dir := path.Dir(part.File)
		slashIndex := strings.LastIndex(dir, "/")
		if strings.Contains(dir[slashIndex+1:], "-") {
			continue
		}
		yearVal, quarterVal := dir[slashIndex+1:][:4], dir[slashIndex+1:][4:]

		if year != "" && year != yearVal {
			continue
		}
		if quarter != "" && quarter != quarterVal {
			continue
		}
		filteredParts = append(filteredParts, part)
	}

	return filteredParts
}

func (parts Partitions) Size() float64 {
	var size float64

	for _, part := range parts {
		size += part.Size
	}

	return size
}
