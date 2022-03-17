package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type alien struct {
	Name       string   `json:"name"`
	ImgURL     string   `json:"imgURL"`
	Species    string   `json:"species"`
	HomePlanet string   `json:"homePlanet"`
	Powers     []string `json:"powers"`
}

func main() {
	classicAliens := getAliens("https://ben10.fandom.com/wiki/Category:Original_Series_Aliens")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		enc.Encode(classicAliens)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
