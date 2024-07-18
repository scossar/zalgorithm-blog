package utils

import (
	"path/filepath"
	"strings"
)

func FilesOfType(dir string, ext string) ([]string, error) {
	pattern := filepath.Join(dir, "*."+ext)
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}

type FileInfo struct {
	Name  string
	Title string
	Path  string
}

func Info(files []string) []FileInfo {
	var fileInfos []FileInfo
	for _, file := range files {
		name := filepath.Base(file)
		title := strings.TrimSuffix(name, filepath.Ext(name))
		fileInfos = append(fileInfos, FileInfo{
			Name:  name,
			Title: title,
			Path:  file,
		})
	}
	return fileInfos
}

type TitleAndSlug struct {
	Title string
	Slug  string
}

func TitlesAndSlugs(files []string) []TitleAndSlug {
	var titlesAndSlugs []TitleAndSlug
	for _, file := range files {
		name := filepath.Base(file)
		title := strings.TrimSuffix(name, filepath.Ext(name))
		slug := strings.ReplaceAll(title, " ", "-")

		titlesAndSlugs = append(titlesAndSlugs, TitleAndSlug{
			Title: title,
			Slug:  slug,
		})
	}
	return titlesAndSlugs
}
