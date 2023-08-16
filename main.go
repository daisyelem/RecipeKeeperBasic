package main

import (
	"log"
	"net/http"

	"github.com/daisyelem/RecipeKeeperBasic/data"
	"github.com/daisyelem/RecipeKeeperBasic/handlers"
)

func main() {
	data.FetchAllRecipes()
	//fmt.Println(data.AllRecipes)

	server := http.NewServeMux()

	style := http.FileServer(http.Dir("frontend/css"))

	server.Handle("/frontend/css/", http.StripPrefix("/frontend/css/", style))

	server.HandleFunc("/", handlers.HomePage)
	server.HandleFunc("/recipe/", handlers.RecipePage)

	log.Fatal(http.ListenAndServe(":8000", server))
}
