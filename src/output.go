package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	serverAddress  = "localhost:8080"
	outputFilename = "aliens.csv"
)

func handleServer(aliensData []alien) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		enc.Encode(aliensData)
	})

	fmt.Fprintf(os.Stderr, "listening at http://%s\n", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil))
}

func handleFile(aliensData []alien) {
	fileName := "aliens.csv"

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	f, _ := os.Create(fileName)
	defer f.Close()

	fileData := ""

	for _, alien := range aliensData {
		fileData += prepareAlienEntryCSV(alien)
	}

	f.WriteString(fileData)
}

func prepareAlienEntryCSV(a alien) string {
	return fmt.Sprintf("%s;%s;%s;%s;%s\n", a.Name, a.ImgURL, a.Species, a.HomePlanet, strings.Join(a.Powers, ","))
}
