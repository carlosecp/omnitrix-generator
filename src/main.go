package main

import "flag"

const (
	classicAliensWikiURL    = "https://ben10.fandom.com/wiki/Category:Original_Series_Aliens"
	alienForceAliensWikiURL = "https://ben10.fandom.com/wiki/Category:Alien_Force_Aliens"
	ultimeAliensWikiURL     = "https://ben10.fandom.com/wiki/Category:Ultimate_Alien_Aliens"
)

var startServer = flag.Bool("server", false, "dump the data to a web server instead of saving it to a file")

func main() {
	flag.Parse()

	wikisURLs := []string{
		classicAliensWikiURL,
		alienForceAliensWikiURL,
		ultimeAliensWikiURL,
	}

	run(*startServer, wikisURLs)
}

func run(startServer bool, wikisURLs []string) {
	aliensData := []alien{}

	for _, wikiURL := range wikisURLs {
		aliensData = append(aliensData, getAliensData(wikiURL)...)
	}

	if startServer {
		handleServer(aliensData)
	} else {
		handleFile(aliensData)
	}
}
