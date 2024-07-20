package main

import (
	"log"
	"net/http"

	"github.com/scossar/zalgorithm-blog/db"
	"github.com/scossar/zalgorithm-blog/handlers"
)

func main() {
	db.PrepareDB("/home/scossar/obsidian_vault")
	indexHandler := handlers.NewIndexHandler()
	noteHandler := handlers.NewNoteHandler()

	http.HandleFunc("/", indexHandler.Handler)
	http.HandleFunc("/note/{slug}/{id}", noteHandler.NoteHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
