package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"github.com/scossar/zalgorithm-blog/utils"
)

type NoteHandler struct{}

func NewNoteHandler() *NoteHandler {
	return &NoteHandler{}
}

func (h *NoteHandler) NoteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatalf("Error opening database file: %v", err)
	}

	defer db.Close()

	noteQuery := `SELECT title, markdown FROM notes WHERE id = ?`
	stmt, err := db.Prepare(noteQuery)
	if err != nil {
		log.Fatalf("Error preparing noteQuery: %v", err)
	}

	var title string
	var markdown string
	err = stmt.QueryRow(id).Scan(&title, &markdown)
	if err != nil {
		log.Fatalf("Error querying db: %v", err)
	}

	defer stmt.Close()

	type Note struct {
		Title string
		HTML  template.HTML
	}

	html := utils.MDToHTML([]byte(markdown))
	note := Note{Title: title, HTML: template.HTML(html)}
	t, err := template.ParseFiles("templates/layout.html", "templates/note.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	exErr := t.Execute(w, note)
	if exErr != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
