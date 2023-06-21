package http_test

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "01.kood.tech/git/AmrKharaba/ascii-art/http"
	pages "01.kood.tech/git/AmrKharaba/ascii-art/templates"
)

func AssertEqual[T comparable](t *testing.T, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("\ngot %v;\n\n\n\n\n\n\n expected %v\n\n", actual, expected)
	}
}

func TestGETHandler(t *testing.T) {
	// t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handlers.GET(rec, req)
	res := rec.Body

	expectedTemplate, err := fs.ReadFile(pages.BaseHTML, "base.html")
	if err != nil {
		t.Fatal(err)
	}
	expected := string(expectedTemplate)
	actual := res.String()

	AssertEqual[string](t, actual, expected)

}

func TestPOSTHandler(t *testing.T) {
	// t.Parallel()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()

	handlers.GET(rec, req)
	res := rec.Body

	expectedTemplate, err := fs.ReadFile(pages.AsciiArtHTML, "ascii-art.html")
	if err != nil {
		t.Fatal(err)
	}
	expected := string(expectedTemplate)
	actual := res.String()

	AssertEqual[string](t, actual, expected)

}

func TestRoutes(t *testing.T) {
	server := httptest.NewServer(handlers.Routes())
	defer server.Close()
}
