package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/daisyelem/RecipeKeeperBasic/data"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("frontend/homePage.html")
	if err != nil {
		log.Fatal(err, "ERROR: problem with file path")
	}

	tmpl.Execute(w, data.AllRecipes)

}
