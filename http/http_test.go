package http_test

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	handlers "01.kood.tech/git/AmrKharaba/ascii-art/http"
	pages "01.kood.tech/git/AmrKharaba/ascii-art/templates"
)

// AssertEqual is an helper test function to check equality of `got` and `expected`
func AssertEqual[T comparable](t *testing.T, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("--GOT--: %v --EXPECTED--: %v", got, expected)
	}
}

// AssertContains is an helper test function to each if `got` contains an expected string `ascii-art`
func AssertContains[T comparable](t *testing.T, got string) {
	t.Helper()

	if !strings.Contains(got, "ascii-art") {
		t.Error("\ncorrect template not loaded, InjectedString not dectected\n")
	}
}

// TestGET tests that the GET handlers renders the expected template
func TestGET(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handlers.GET(rec, req)
	res := rec.Body

	expectedTemplate, err := fs.ReadFile(pages.BaseHTML, "base.html")
	if err != nil {
		t.Fatal(err)
	}
	expected := string(expectedTemplate)
	got := res.String()

	AssertEqual[string](t, got, expected)
}

// TestPOST tests that the GET handlers renders the expected template
func TestPOST(t *testing.T) {

	tests := map[string]struct {
		name  string
		style string
		query string
	}{
		"standard":   {name: "standard", style: "standard.txt", query: "STANDARD"},
		"shadow":     {name: "shadow", style: "shadow.txt", query: "SHADOW"},
		"thinkertoy": {name: "thinkertoy", style: "thinkertoy.txt", query: "THINKERTOY"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			formData := url.Values{}
			formData.Set("style", tc.style)
			formData.Set("query", tc.query)

			req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(formData.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			rec := httptest.NewRecorder()

			handlers.POST(rec, req)
			res := rec.Body

			got := res.String()

			AssertContains[string](t, got)
		})
	}

}

// TestRoutesAndHandler tests that the right url route invokes the correct handler func
func TestRoutesAndHandler(t *testing.T) {

	tests := map[string]struct {
		handler http.Handler
		route   string
		body    io.Reader
		style   string
		query   string
	}{
		"GET":  {handler: http.HandlerFunc(handlers.GET), route: "/", body: nil},
		"POST": {handler: http.HandlerFunc(handlers.POST), route: "/ascii-art", body: nil, style: "standard.txt", query: "STANDARD"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {

			server := httptest.NewServer(tc.handler)
			defer server.Close()

			var (
				resp *http.Response
				err  error
			)

			formData := url.Values{}
			formData.Set("style", tc.style)
			formData.Set("query", tc.query)

			switch name {
			case "GET":
				resp, err = http.Get(server.URL)
			case "POST":
				resp, err = http.Post(fmt.Sprintf("%s%s", server.URL, tc.route), "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
			default:
				t.Fatalf("the %v was not key in tests map for TestRoutesAndHandler", name)
			}

			if err != nil {
				t.Errorf("could not send %s request; error is %v", name, err)
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("\nexpected status OK; got %s\n", resp.Status)
			}
		})
	}
}

// TestRecover tests the Recover middleware
func TestRecover(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handlers.Recover(next).ServeHTTP(res, req)

	resp := res.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("\nexpected status OK; got %s\n", resp.Status)
	}

	AssertEqual[string](t, res.Body.String(), "OK")
}

// TestLogRequest tests the LogRequest middleware
func TestLogRequest(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handlers.LogRequest(next).ServeHTTP(res, req)

	resp := res.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("\nexpected status OK; got %s\n", resp.Status)
	}

	AssertEqual[string](t, res.Body.String(), "OK")
}

// TestHttpMethod tests the HttpMethod middleware
func TestHttpMethod(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handlers.HttpMethod(http.MethodGet, next).ServeHTTP(res, req)

	resp := res.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("\nexpected status OK; got %s\n", resp.Status)
	}

	AssertEqual[string](t, res.Body.String(), "OK")
}

// TestGETRoute tests the `/` url route
func TestGETRoute(t *testing.T) {
	server := httptest.NewServer(handlers.Routes())
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("could not send %s request; error is %v", "/", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("\nexpected status OK; got %s\n", resp.Status)
	}
}

// TestPOSTRoute tests the `/ascii-art` url route
func TestPOSTRoute(t *testing.T) {
	server := httptest.NewServer(handlers.Routes())
	defer server.Close()

	formData := url.Values{}
	formData.Set("style", "standard.txt")
	formData.Set("query", "STANDARD")

	resp, err := http.Post(fmt.Sprintf("%s%s", server.URL, "/ascii-art"), "application/x-www-form-urlencoded", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Errorf("could not send %s request; error is %v", "/", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("\nexpected status OK; got %s\n", resp.Status)
	}
}
