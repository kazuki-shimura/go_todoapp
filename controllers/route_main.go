package controllers

import (
	"log"
	"html/template"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/templates/top.html")
	if err != nil {
		log.Fatalln()
	}
	t.Execute(w, "Hiiiiiii")
}