package main

import (
	"fmt"
	"log"
	"net/http"
	"recipes/data"
	"recipes/handlers"
)

func main() {
	data.FetchAllRecipes()

	server := http.NewServeMux()

	style := http.FileServer(http.Dir("frontend/css"))

	server.Handle("/frontend/css/", http.StripPrefix("/frontend/css/", style))

	server.HandleFunc("/", handlers.HomePage)
	server.HandleFunc("/recipe/", handlers.RecipePage)

	fmt.Println(data.AllRecipes)

	log.Fatal(http.ListenAndServe(":8000", server))
}
