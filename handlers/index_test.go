package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/scossar/zalgorithm-blog/handlers"
)

type MockFileFetcher struct{}

func (m MockFileFetcher) FilesOfType(dir, fileType string) ([]string, error) {
	return []string{"mockfile one.md", "mockfile two.md"}, nil
}

func TestMain(m *testing.M) {
	// needed to load the templates
	if err := os.Chdir("../"); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestIndexHandler(t *testing.T) {
	mockFetcher := MockFileFetcher{}
	handler := handlers.NewHandler(mockFetcher, "/mock/dir")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.IndexHandler).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedSubstring := "mockfile one"
	if !strings.Contains(rr.Body.String(), expectedSubstring) {
		t.Errorf("expected body to contain %v got %v", expectedSubstring, rr.Body.String())
	}
}
