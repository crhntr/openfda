package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
		}
	}
}

func eachFile(fun func(f *os.File, fname string) error) {
	downloads := OpenDownloads()

	parts := downloads.Results.Drug.Event.Partitions.Filter(*year, *quarter)

	for i, part := range parts {
		log.Print(part.Name)
		func(i int, part Partition) {
			dir := path.Dir(strings.TrimPrefix(part.File, "https://download.open.fda.gov/"))
			dir = path.Join(*outPath, dir)

			files, _ := ioutil.ReadDir(dir)
			for _, file := range files {
				if file.Name() == ".DS_Store" {
					continue
				}
				f, err := os.Open(path.Join(dir, file.Name()))
				if err != nil {
					log.Printf("  %d %s %s", i, file.Name(), err)
					continue
				}
				func() {
					defer f.Close()
					if err := fun(f, part.File); err != nil {
						log.Printf("  %q in %q", err, part.File)
					}
				}()
			}
		}(i, part)
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
