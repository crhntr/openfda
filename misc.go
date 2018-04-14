package main

import "fmt"

func calcSize() {
	counts := map[string]int{}
	downloads := OpenDownloads()
	fmt.Printf("parts: %d\n", len(downloads.Results.Drug.Event.Partitions))

	for _, part := range downloads.Results.Drug.Event.Partitions {
		_, pat := ShiftPath(part.File)
		_, pat = ShiftPath(pat) // drug
		_, pat = ShiftPath(pat) // event
		_, pat = ShiftPath(pat) // event
		qt, _ := ShiftPath(pat)
		counts[qt] += part.Records
		fmt.Printf("%10d %10f %s %s\n", part.Records, part.Size, qt, part.File)
	}

	for key, val := range counts {
		fmt.Printf("%s: %d\n", key, val)
	}
}
