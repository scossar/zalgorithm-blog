package main

import (
	"log"

	"github.com/scossar/zalgorithm-blog/utils"
)

func main() {
	files, err := utils.FilesOfType("/home/scossar/obsidian_vault", "md")
	if err != nil {
		log.Fatalf("An error was returned from the call to FilesOfType: %v", err)
	}

	infos := utils.Info(files)

	for _, info := range infos {
		log.Printf("name: %v", info.Name)
		log.Printf("title: %v", info.Title)
		log.Printf("path: %v", info.Path)
	}
}
