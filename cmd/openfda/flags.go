package main

import "flag"

var (
	outPath = flag.String("out", "openfda_data", "Output Folder for Downloads")
	year    = flag.String("year", "", "year of files to download")
	quarter = flag.String("quarter", "", "year of files to download")

	limit = flag.Int("limit", 100, "limit something")
)
