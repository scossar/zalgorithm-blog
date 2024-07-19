package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scossar/zalgorithm-blog/handlers"
	"github.com/scossar/zalgorithm-blog/utils"
)

func main() {
	realFileFetcher := utils.RealFileFetcher{}
	notesDir := "/home/scossar/obsidian_vault"
	handler := handlers.NewIndexHandler(realFileFetcher, notesDir)
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler)

	noteHandler := handlers.NewNoteHandler(notesDir)
	r.HandleFunc("/note/{slug}", noteHandler.NoteHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
