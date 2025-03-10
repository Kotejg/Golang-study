package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola Mundo"))
}

type user struct {
	Nome string
}

func main() {

	templates = template.Must(templates.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	usuario := user{"Jeff"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", usuario)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
