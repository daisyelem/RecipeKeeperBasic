package handlers

import (
	"net/http"
	"html/template"
	"log"
	"recipes/data"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/homePage.html")
	if err != nil {
		log.Fatal(err, "ERROR: problem with file path")
	}

	tmpl.Execute(w, data.AllRecipes)

}
