package http

import (
	"fmt"
	"net/http"
	"text/template"

	domain "01.kood.tech/git/AmrKharaba/ascii-art"
	pages "01.kood.tech/git/AmrKharaba/ascii-art/templates"
)

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

	err = ts.Execute(w, domain.ProduceAsciiArt(text, style))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err)
	}

}
