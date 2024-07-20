package db

import (
	"database/sql"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/scossar/zalgorithm-blog/utils"
)

func PrepareDB(rootDir string) {
	db, err := sql.Open("sqlite3", "./notes.db")
	checkErr((err))
	defer db.Close()

	ensureTablesExist(db)

	seed(db, rootDir)

	// rows, err := db.Query("SELECT id, title, slug, markdown FROM notes")
	// checkErr(err)
	//
	// defer rows.Close()
	//
	// for rows.Next() {
	// 	var id int
	// 	var title, slug, markdown string
	// 	err = rows.Scan(&id, &title, &slug, &markdown)
	// 	checkErr(err)
	// 	fmt.Printf("ID: %d\nTitle: %s\nSlug: %s\nMarkdown: %s\n", id, title, slug, markdown)
	// }
	//
	// if err = rows.Err(); err != nil {
	// 	panic(err)
	// }
}

func seed(db *sql.DB, rootDir string) {
	var version int
	err := db.QueryRow("SELECT version FROM version WHERE id = 1").Scan(&version)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	if version == 1 {
		return
	}

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
		insert(db, title, slug, md)
	}

	_, err = db.Exec("INSERT OR REPLACE INTO version (id, version) VALUES (1, 1)")
	checkErr(err)
}

func ensureTablesExist(db *sql.DB) {
	createTableQuery := `
  CREATE TABLE IF NOT EXISTS notes (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      title TEXT NOT NULL,
      slug TEXT NOT NULL,
      markdown TEXT NOT NULL
  );
  CREATE TABLE IF NOT EXISTS version (
      id INTEGER PRIMARY KEY,
      version INTEGER NOT NULL
  )
  `
	_, err := db.Exec(createTableQuery)
	checkErr(err)
}

func insert(db *sql.DB, title, slug, md string) {
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
