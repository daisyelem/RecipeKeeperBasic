package data

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FetchAllRecipes() {
	jsonInfo, err := http.Get("https://api.npoint.io/fff0f131782057b16a12")

	if err != nil {
		log.Fatal(err, "ERROR: problem with url")
	}

	defer jsonInfo.Body.Close()

	body, err := io.ReadAll(jsonInfo.Body)

	if err != nil {
		log.Fatal(err, "ERROR: problem reading data")
	}

	json.Unmarshal(body, &AllRecipes)
}
