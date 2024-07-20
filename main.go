package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scossar/zalgorithm-blog/db"
	"github.com/scossar/zalgorithm-blog/handlers"
)

func main() {
	db.PrepareDB("/home/scossar/obsidian_vault")
	handler := handlers.NewIndexHandler()
	// TODO: I don't think gorilla/mux is needed to deal with paths
	r := mux.NewRouter()

	r.HandleFunc("/", handler.IndexHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
