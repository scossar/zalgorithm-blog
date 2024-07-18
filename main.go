package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scossar/zalgorithm-blog/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.IndexHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
