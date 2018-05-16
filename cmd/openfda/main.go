package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	apiURL = "https://api.fda.gov"

	downloadsFilePath = apiURL + "/download.json"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	default:

		fmt.Println("unknown argument")
	case "create-test-data":
		createTestData("patient.drug.openfda.brand_name:%22kalydeco%22+AND+receivedate:[20170901+TO+20171231]")
	case "fetch-download-file":
		fetchDownloadsFile()

	case "drug":
		switch flag.Arg(1) {

		case "event":
			switch flag.Arg(2) {

			case "size":
				downloads := openDownloads()
				fmt.Println(downloads.Results.Drug.Event.Partitions.FilterDrugEvents(*year, *quarter).Size())

			case "import":
				importFiles()

			case "import-all":
				importAll()

			case "count":
				countDrugEvents()

			case "count-files":
				var count int
				eachDrugEventFile(func(f *os.File) error {
					count++
					return nil
				})
				fmt.Printf("\n\nfound %d files\n\n", count)

			case "ls":
				eachDrugEventFile(func(f *os.File) error {
					fmt.Println(f.Name())
					return nil
				})

			}

		case "label":

		}
	}
}
