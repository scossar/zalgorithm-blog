package utils

import (
	"os"
	"path/filepath"
	"reflect"
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

func TestInfo(t *testing.T) {
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
		t.Fatalf("Files of type returned an error: %v", err)
	}

	info := Info(files)

	if expectedLen := 2; len(info) != expectedLen {
		t.Errorf("Expected info on %d markdown files, got %d", expectedLen, len(info))
	}

	expectedNames := map[string]bool{"test1.md": true, "test2.md": true}
	for _, fileInfo := range info {
		if _, ok := expectedNames[fileInfo.Name]; !ok {
			t.Errorf("Unexpected file info: %+v", fileInfo)
		}
	}
}

// keeping this as an example of using reflect.DeepEqual to compare structs
// the previous test is probably all that's needed though.
func TestInfoDeepEqual(t *testing.T) {
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
		t.Fatalf("Files of type returned an error: %v", err)
	}

	info := Info(files)

	if expectedLen := 2; len(info) != expectedLen {
		t.Errorf("Expected info on %d markdown files, got %d", expectedLen, len(info))
	}

	expectedFileInfos := []FileInfo{
		{Name: "test1.md", Title: "test1", Path: filepath.Join(dir, "test1.md")},
		{Name: "test2.md", Title: "test2", Path: filepath.Join(dir, "test2.md")},
	}

	for _, expected := range expectedFileInfos {
		found := false
		for _, actual := range info {
			if reflect.DeepEqual(expected, actual) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file info not found: %+v", expected)
		}
	}
}
