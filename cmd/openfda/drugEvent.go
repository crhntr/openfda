package main

import (
	"log"
	"os"
	"path"
)

func eachDrugEventFile(fun func(f *os.File) error) {
	downloads := openDownloads()

	parts := downloads.Results.Drug.Event.Partitions.Filter(*year, *quarter)

	for i, p := range parts {
		func(i int, part Partition) {
			filename := path.Join(*outPath, part.File[len("https://download.open.fda.gov/"):len(part.File)-len(".zip")])
			f, err := os.Open(filename)
			if err != nil {
				log.Printf("  %d %s %s", i, filename, err)
				return
			}
			func() {
				defer f.Close()
				if err := fun(f); err != nil {
					log.Printf("  %q in %q", err, filename)
				}
			}()
		}(i, p)
	}
}
