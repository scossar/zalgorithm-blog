package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/scossar/zalgorithm-blog/utils"
)

type Handler struct {
	FileFetcher utils.FileFetcher
	NotesDir    string
}

func NewIndexHandler(fileFetcher utils.FileFetcher, notesDir string) *Handler {
	return &Handler{FileFetcher: fileFetcher, NotesDir: notesDir}
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	// mdFiles, err := utils.FilesOfType(notesDir, "md")
	mdFiles, err := h.FileFetcher.FilesOfType(h.NotesDir, "md")
	if err != nil {
		log.Fatalf("Error returning markdown files: %v", err)
	}

	titlesAndSlugs := utils.TitlesAndSlugs(mdFiles)

	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = t.Execute(w, titlesAndSlugs)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
