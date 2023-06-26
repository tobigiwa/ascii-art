package main

import (
	"fmt"
	"log"
	"net/http"

	web "01.kood.tech/git/AmrKharaba/ascii-art/http"
)

func main() {
	// domain.ProduceAsciiArt()

	webserver := &http.Server{
		Addr:    ":5030",
		Handler: web.Routes(),
		// ErrorLog: ,
	}

	fmt.Println("server running...@ http://localhost:5030")
	log.Fatal(webserver.ListenAndServe())
}
