package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/scossar/zalgorithm-blog/utils"
)

const notesDir = "/home/scossar/obsidian_vault"

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	mdFiles, err := utils.FilesOfType(notesDir, "md")
	if err != nil {
		log.Fatalf("Error returning markdown files: %v", err)
	}

	// TODO: probably all the data returned from Info isn't required here
	info := utils.Info(mdFiles)

	t, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = t.Execute(w, info)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
