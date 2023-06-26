package http

import (
	"fmt"
	"net/http"
	"text/template"

	"01.kood.tech/git/AmrKharaba/ascii-art/domain"
	pages "01.kood.tech/git/AmrKharaba/ascii-art/templates"
)

// GET is the handler that returns the base.hmtl template
func GET(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFS(pages.BaseHTML, "base.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(fmt.Errorf("could not found base.html"))
	}

	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err)
	}
}

// POST is the handler that returns the template containing the Ascii-Art
func POST(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	var (
		text, style string
	)

	style = r.PostForm.Get("style")
	if text = r.PostForm.Get("query"); text == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(fmt.Errorf("no text was entered"))
	}

	ts, err := template.ParseFS(pages.AsciiArtHTML, "ascii-art.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(fmt.Errorf("could not found ascii-art.html"))
	}

	dynamicData := struct {
		// InjectedString is included to aid in testing
		InjectedString string
		// AsciiArt is the generated banner style
		AsciiArt string
	}{
		"ascii-art",
		domain.ProduceAsciiArt(text, style),
	}

	ts.Execute(w, dynamicData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err)
	}

}
