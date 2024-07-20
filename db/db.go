package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/scossar/zalgorithm-blog/utils"
)

func SeedDb(rootDir string) {
	db, err := sql.Open("sqlite3", "./notes.db")
	checkErr((err))

	defer db.Close()

	createTableQuery := `
  CREATE TABLE IF NOT EXISTS notes (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT NOT NULL,
      slug TEXT NOT NULL,
      markdown TEXT NOT NULL
  );
  CREATE TABLE IF NOT EXISTS version (
      id INTEGER PRIMARY KEY,
      version INTEGER PRIMARY KEY
  )
  `
	_, err = db.Exec(createTableQuery)
	checkErr(err)

	mdFiles, err := utils.FilesOfType(rootDir, "md")
	checkErr(err)

	for _, file := range mdFiles {
		mdBytes, err := os.ReadFile(file)
		checkErr(err)
		md := string(mdBytes)
		name := filepath.Base(file)
		title := strings.TrimSuffix(name, filepath.Ext(name))
		slug := strings.ReplaceAll(title, " ", "-")
		re := regexp.MustCompile(`[^\w-]+`)
		slug = re.ReplaceAllString(slug, "")
		slug = strings.ToLower(slug)
		insert(md, title, slug)

	}

	rows, err := db.Query("SELECT id, title, markdown FROM notes")
	checkErr(err)

	defer rows.Close()

	for rows.Next() {
		var id int
		var title, markdown string
		err = rows.Scan(&id, &title, &markdown)
		checkErr(err)
		fmt.Printf("ID: %d\nTitle: %s\nMarkdown: %s\n", id, title, markdown)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
}

func insert(title, slug, md string) {
	db, err := sql.Open("sqlite3", "./notes.db")
	checkErr(err)
	insertNoteQuery := `INSERT INTO notes (title, slug, markdown) VALUES (?, ?, ?)`
	stmt, err := db.Prepare(insertNoteQuery)
	checkErr(err)

	defer stmt.Close()

	_, err = stmt.Exec(title, slug, md)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
