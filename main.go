package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	outPath = *flag.String("out", "data", "Output Folder for Downloads")
)

func main() {
	flag.Parse()

	f, err := os.Open("downloads.json") // https://api.fda.gov/download.json
	if err != nil {
		panic(err)
	}
	var downlaods Downloads
	if err = json.NewDecoder(f).Decode(&downlaods); err != nil {
		panic(err)
	}

	switch flag.Arg(0) {
	case "drug.event":
		switch flag.Arg(1) {
		case "size":
			fmt.Println(downlaods.Results.Drug.Event.Size())
		case "parse":
			// for i, file := range flag.Args()[2:] {
			//
			// }
		case "download":
			size := downlaods.Results.Drug.Event.Size()
			var downlaoded float64

			for i, part := range downlaods.Results.Drug.Event.Partitions {
				func(i int, part Partition) {
					dir := path.Dir(strings.TrimPrefix(part.File, "https://download.open.fda.gov/"))

					if err := os.MkdirAll(path.Join(outPath, dir), os.ModePerm); err != nil {
						log.Printf("%d %s %s", i, part.File, err)
						return
					}

					log.Printf("%d %q %.2f %.2f", i, part.Name, part.Size, (downlaoded/size)*100)
					resp, err := http.Get(part.File)
					if err != nil {
						log.Printf("%d %s %s", i, part.File, err)
						return
					}
					defer resp.Body.Close()

					out, err := os.Create(path.Join(outPath, dir, path.Base(part.File)))
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
	}
}
