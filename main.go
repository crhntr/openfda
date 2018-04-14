package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/crhntr/openfda/drug"
	"github.com/globalsign/mgo"
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
			session, err := mgo.Dial(":27017")
			if err != nil {
				log.Fatal(err)
				return
			}

			for i, fileName := range flag.Args()[2:] {
				if err := func(i int, filename string) error {
					log.Printf("%d %q", i, fileName)

					f, err := os.Open(fileName)
					if err != nil {
						return err
					}
					defer f.Close()

					dec := json.NewDecoder(f)

					for {
						t, err := dec.Token()
						if err == io.EOF {
							return nil
						}
						if err != nil {
							return err
						}
						// fmt.Printf("%T %v\n", t, t)
						if str, ok := t.(string); ok && str == "results" {
							t, err = dec.Token()
							if err != nil {
								return err
							}
							if val, ok := t.(json.Delim); ok && val == json.Delim('[') {
								break
							}
						}
					}

					for {
						var rawEvent drug.RawEvent
						if err := dec.Decode(&rawEvent); err != nil {
							return err
						}

						event, drugData := rawEvent.Event()
						event.FileName = fileName

						if err := session.DB("openfda").C("drug_event").Insert(event); err != nil {
							if !mgo.IsDup(err) {
								log.Printf("%d %s %s %s", i, event.SafetyReportID, filename, err)
							}
						}

						for _, d := range drugData {
							if d.ID != "" {
								if err := session.DB("openfda").C("drug").Insert(d); err != nil {
									if !mgo.IsDup(err) {
										log.Printf("%d %s %s %s", i, event.SafetyReportID, filename, err)
									}
								}
							}
						}

						if !dec.More() {
							return nil
						}
					}
					return nil
				}(i, fileName); err != nil {
					log.Printf("%d %s %s", i, fileName, err)
				}
			}
		case "import-files":
			session, err := mgo.Dial(":27017")
			if err != nil {
				log.Fatal(err)
				return
			}

			for i, fname := range jsonFiles {
				if err := func(i int, filename string) error {
					log.Printf("%d %q", i, filename)

					f, err := os.Open(filename)
					if err != nil {
						return err
					}
					defer f.Close()

					dec := json.NewDecoder(f)

					for {
						t, err := dec.Token()
						if err == io.EOF {
							return nil
						}
						if err != nil {
							return err
						}
						// fmt.Printf("%T %v\n", t, t)
						if str, ok := t.(string); ok && str == "results" {
							t, err = dec.Token()
							if err != nil {
								return err
							}
							if val, ok := t.(json.Delim); ok && val == json.Delim('[') {
								break
							}
						}
					}

					for {
						var rawEvent drug.RawEvent
						if err := dec.Decode(&rawEvent); err != nil {
							return err
						}

						event, drugData := rawEvent.Event()
						event.FileName = filename

						if err := session.DB("openfda").C("drug_event").Insert(event); err != nil {
							if !mgo.IsDup(err) {
								log.Printf("%d %s %s %s", i, event.SafetyReportID, filename, err)
							}
						}

						for _, d := range drugData {
							if d.ID != "" {
								if err := session.DB("openfda").C("drug").Insert(d); err != nil {
									if !mgo.IsDup(err) {
										log.Printf("%d %s %s %s", i, event.SafetyReportID, filename, err)
									}
								}
							}
						}

						if !dec.More() {
							return nil
						}
					}
					return nil
				}(i, fname); err != nil {
					log.Printf("%d %s %s", i, fname, err)
				}
			}
		case "download":
			downloads := OpenDownloads()
			size := downloads.Results.Drug.Event.Size()
			var downlaoded float64

			parts := downloads.Results.Drug.Event.Partitions.Filter(*year, *quarter)
			fmt.Println(len(parts))
			for i, part := range parts {
				func(i int, part Partition) {
					dir := path.Dir(strings.TrimPrefix(part.File, "https://download.open.fda.gov/"))

					if err := os.MkdirAll(path.Join(*outPath, dir), os.ModePerm); err != nil {
						log.Printf("%d %s %s", i, part.File, err)
						return
					}

					log.Printf("  %d %q %.2f %.2f", i, part.Name, part.Size, (downlaoded/size)*100)
					resp, err := http.Get(part.File)
					if err != nil {
						log.Printf("  %d %s %s", i, part.File, err)
						return
					}
					defer resp.Body.Close()

					out, err := os.Create(path.Join(*outPath, dir, path.Base(part.File)))
					if err != nil {
						log.Printf("%d %s %s", i, part.File, err)
						return
					}
					defer out.Close()
					io.Copy(out, resp.Body)
					downlaoded += part.Size
				}(i, part)
			}
		case "import-all":

			session, err := mgo.Dial(":27017")
			if err != nil {
				log.Fatal(err)
				return
			}

			eachFile(func(f *os.File, fileName string) error {
				log.Printf("-> %q", fileName)

				dec := json.NewDecoder(f)

				for {
					t, err := dec.Token()
					if err == io.EOF {
						return nil
					}
					if err != nil {
						return err
					}
					// fmt.Printf("%T %v\n", t, t)
					if str, ok := t.(string); ok && str == "results" {
						t, err = dec.Token()
						if err != nil {
							return err
						}
						if val, ok := t.(json.Delim); ok && val == json.Delim('[') {
							break
						}
					}
				}

				for {
					var rawEvent drug.RawEvent
					if err := dec.Decode(&rawEvent); err != nil {
						return err
					}

					event, drugData := rawEvent.Event()
					event.FileName = fileName

					if err := session.DB("openfda").C("drug_event").Insert(event); err != nil {
						if !mgo.IsDup(err) {
							log.Printf("  %s %s %s", event.SafetyReportID, fileName, err)
						}
					}

					for _, d := range drugData {
						if d.ID != "" {
							if err := session.DB("openfda").C("drug").Insert(d); err != nil {
								if !mgo.IsDup(err) {
									log.Printf("  %s %s %s", event.SafetyReportID, fileName, err)
								}
							}
						}
					}

					if !dec.More() {
						return nil
					}
				}
				return nil
			})
		case "stats":
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

func OpenDownloads() Downloads {
	f, err := os.Open("downloads.json") // https://api.fda.gov/download.json
	if err != nil {
		panic(err)
	}
	var downloads Downloads
	if err = json.NewDecoder(f).Decode(&downloads); err != nil {
		panic(err)
	}
	return downloads
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
