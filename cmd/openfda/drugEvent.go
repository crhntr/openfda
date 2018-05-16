package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func eachDrugEventFile(fun func(f *os.File) error) {
	downloads := openDownloads()

	parts := downloads.Results.Drug.Event.Partitions.Filter(*year, *quarter)

	for i, p := range parts {
		func(i int, part Partition) {
			filename := path.Join(*outPath, part.File[len(downloadsFilePath):len(part.File)-len(".zip")])
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

func countDrugEvents() {
	counts := map[string]int{}
	downloads := openDownloads()
	fmt.Printf("parts: %d\n", len(downloads.Results.Drug.Event.Partitions))

	for _, part := range downloads.Results.Drug.Event.Partitions {
		_, pat := ShiftPath(part.File)
		_, pat = ShiftPath(pat) // drug
		_, pat = ShiftPath(pat) // event
		_, pat = ShiftPath(pat) // event
		qt, _ := ShiftPath(pat)
		counts[qt] += part.Records
		// fmt.Printf("%10d %10f %s %s\n", part.Records, part.Size, qt, part.File)
	}

	for key, val := range counts {
		fmt.Printf("%s: %d\n", key, val)
	}
}
