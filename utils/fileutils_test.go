package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilesOfType(t *testing.T) {
	dir := t.TempDir()
	_, err := os.Create(filepath.Join(dir, "test1.md"))
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err = os.Create(filepath.Join(dir, "test2.md"))
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err = os.Create(filepath.Join(dir, "test3.txt"))
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	files, err := FilesOfType(dir, "md")
	if err != nil {
		t.Fatalf("FilesOfType returned an error: %v", err)
	}

	if expected := 2; len(files) != expected {
		t.Errorf("Expected %d markdown files, got %d", expected, len(files))
	}
}
