package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type Handler struct{}

func NewIndexHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./notes.db")
	if err != nil {
		log.Fatalf("Error opening database file: %v", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, slug FROM notes")
	if err != nil {
		log.Fatalf("Error obtaining rows from db: %v", err)
	}
	defer rows.Close()

	type Note struct {
		Title string
		Slug  string
		ID    int
	}

	var notes []Note

	for rows.Next() {
		var id int
		var title, slug string
		err := rows.Scan(&id, &title, &slug)
		if err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		notes = append(notes, Note{
			ID:    id,
			Title: title,
			Slug:  slug,
		})
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("Error scanning rows: %v", err)
	}

	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = t.Execute(w, notes)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
