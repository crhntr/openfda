package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	default:
		fmt.Println("unknown argument")
	case "downloads-file":
		downloadsFile()
	case "drug":
		switch flag.Arg(1) {
		case "event":
			switch flag.Arg(2) {
			case "size":
				downloads := openDownloads()
				fmt.Println(downloads.Results.Drug.Event.Partitions.Filter(*year, *quarter).Size())
			case "import":
				importFiles()
			case "download":
				download()
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
