package main

import (
	"database/sql"
	"fmt"
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
	dirPaths := []string{*outPath,
		path.Join(*outPath, "drug"),
		path.Join(*outPath, "drug", "event"),
		path.Join(*outPath, "drug", "label"),
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

func showTables(db *sql.DB) error {
	var tables []string
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return fmt.Errorf("could not load tables: %q", err)
	}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		tables = append(tables, name)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating through rows: %q", err)
	}
	fmt.Println(tables)
	return nil
}
