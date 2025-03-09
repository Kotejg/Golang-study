package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola Mundo"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando Usuarios"))
}
func main() {
	http.HandleFunc("/home", home)

	http.HandleFunc("/usuario", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
