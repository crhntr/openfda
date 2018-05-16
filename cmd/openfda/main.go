package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

var (
	outPath = flag.String("out", "data", "Output Folder for Downloads")
	year    = flag.String("year", "", "year of files to download")
	quarter = flag.String("quarter", "", "year of files to download")
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "drug.event":
		switch flag.Arg(1) {
		case "size":
			downloads := OpenDownloads()
			fmt.Println(downloads.Results.Drug.Event.Partitions.Filter(*year, *quarter).Size())
		case "import":
			importFiles()
		case "download":
			download()
		case "import-all":
			importAll()
		case "calc":
			calcSize()
		case "count-files":
			var count int
			eachFile(func(f *os.File) error {
				count++
				return nil
			})
			fmt.Printf("\n\nfound %d files\n\n", count)
		case "ls":
			eachFile(func(f *os.File) error {
				fmt.Println(f.Name())
				return nil
			})
		}
	}
}

func eachFile(fun func(f *os.File) error) {
	downloads := OpenDownloads()

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

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
