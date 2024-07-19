package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/scossar/zalgorithm-blog/utils"
)

type NoteHandler struct {
	NotesDir string
}

func NewNoteHandler(notesDir string) *NoteHandler {
	return &NoteHandler{NotesDir: notesDir}
}

func (h *NoteHandler) NoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	fmt.Printf("Slug: %v", slug)
	// TODO: tmp workaround
	title := strings.ReplaceAll(slug, "-", " ")
	filename := title + ".md"
	filepath := filepath.Join(h.NotesDir, filename)
	html := utils.MdFileToHTML(filepath)

	type Note struct {
		Title string
		Html  template.HTML
	}

	note := Note{Title: title, Html: template.HTML(html)}
	t, err := template.ParseFiles("templates/layout.html", "templates/note.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	exErr := t.Execute(w, note)
	if exErr != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
