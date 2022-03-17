package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

const (
	ben10FandomDomain1 = "ben10.fandom.com"
	ben10FandomDomain2 = "www.ben10.fandom.com"
	cacheDir           = "./ben10_fandom_cache"
)

func getAliens(aliensListURL string) []alien {
	c := colly.NewCollector(
		colly.AllowedDomains(ben10FandomDomain1, ben10FandomDomain2),
		colly.CacheDir(cacheDir),
	)

	c.OnHTML(".category-page__member-link", func(e *colly.HTMLElement) {
		wikiURL := e.Attr("href")
		e.Request.Visit(wikiURL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Fprintf(os.Stderr, "visiting: %s\n", r.URL.String())
	})

	aliens := []alien{}

	c.OnHTML(".portable-infobox", func(e *colly.HTMLElement) {
		name := e.ChildText("[data-source=name]")
		imgURL := e.ChildAttr("img", "src")
		species := e.ChildText("[data-source=species] a")
		homePlanet := e.ChildText("[data-source=species] a")
		powers := []string{}

		e.ForEach("[data-source=power] div", func(_ int, el *colly.HTMLElement) {
			powers = append(powers, getAlienPowers(el)...)
		})

		newAlien := alien{
			Name:       name,
			ImgURL:     imgURL,
			Species:    species,
			HomePlanet: homePlanet,
			Powers:     powers,
		}

		aliens = append(aliens, newAlien)
	})

	c.Visit(aliensListURL)

	return aliens
}

func getAlienPowers(e *colly.HTMLElement) []string {
	powers := []string{}

	htmlContent, err := e.DOM.Html()
	if err != nil {
		return powers
	}

	content := removeHTMLTag(htmlBTag, htmlContent)

	powers = strings.Split(content, htmlBrTag)
	powers = removeEmpty(powers)

	for i, power := range powers {
		if isSurroundedByTag(htmlATag, power) {
			powers[i] = removeHTMLTag(htmlATagContent, power)
		}

		if isSurroundedByTag(htmlSmallTag, power) {
			powers[i] = removeHTMLTag(htmlSmallTagContent, power)
		}

		powers[i] = strings.Trim(powers[i], " ")
	}

	return powers
}
