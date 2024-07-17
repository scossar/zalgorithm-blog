package utils

import "path/filepath"

func FilesOfType(dir string, ext string) ([]string, error) {
	pattern := filepath.Join(dir, "*."+ext)
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}
