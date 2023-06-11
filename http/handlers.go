package http

import (
	"net/http"
	"os"
	"text/template"

	domain "01.kood.tech/git/AmrKharaba/ascii-art"
)

func templateDir() map[string]string {
	path, err := os.Getwd()
	if err != nil {
		panic(err) // panics would be handled and log by the Recover middleware
	}
	return map[string]string{"ascii-art": path + `/templates/ascii-art.html`, "base": path + `/templates/base.html`}

}

func GET(w http.ResponseWriter, r *http.Request) {

	ts, err := template.ParseFiles(templateDir()["base"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err)
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

	var text string
	if text = r.PostForm.Get("query"); text == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		panic(err)
	}

	ts, err := template.ParseFiles(templateDir()["ascii-art"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err)
	}

	err = ts.Execute(w, domain.ProduceAsciiArt(text))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		panic(err)
	}

}
