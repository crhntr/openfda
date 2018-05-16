package main

import (
	"log"
	"os"
	"path"
	"strings"
)

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func joinURL(segments ...string) string {
	return strings.Join(segments, "/")
}

func ensureDataDir() {
	dirPaths := []string{dataDir,
		path.Join(dataDir, "drug"),
		path.Join(dataDir, "drug", "event"),
		path.Join(dataDir, "drug", "label"),
	}

	for _, pth := range dirPaths {
		_, err := os.Stat(pth)
		if os.IsNotExist(err) {
			err := os.Mkdir(pth, 0700)
			if err != nil {
				log.Fatalf("could not create data directory %s: %q", pth, err)
			}
		}
	}
}
