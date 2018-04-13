package main

import (
	"path"
	"strings"
)

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
