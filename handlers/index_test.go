package handlers

import (
	"os"
	"path/filepath"
	"testing"
)

var testDir string

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setup() {
	testDir = filepath.Join(os.TempDir(), "testdir")
	os.MkdirAll(testDir, os.ModePerm)

	os.Create(filepath.Join(testDir, "test one.md"))
	os.Create(filepath.Join(testDir, "test two.md"))
	os.Create(filepath.Join(testDir, "test three.txt"))
}

func tearDown() {
	os.RemoveAll(testDir)
}
