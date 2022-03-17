package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type alien struct {
	Name       string
	ImgURL     string
	Species    string
	HomePlanet string
	Powers     []string
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
