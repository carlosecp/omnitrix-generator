package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func getAliens(aliensListURL string) []alien {
	c := colly.NewCollector(
		colly.AllowedDomains("ben10.fandom.com", "www.ben10.fandom.com"),
		colly.CacheDir("./ben10_fandom_cache"),
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
			name:       name,
			imgURL:     imgURL,
			species:    species,
			homePlanet: homePlanet,
			powers:     powers,
		}

		aliens = append(aliens, newAlien)
	})

	c.Visit(aliensListURL)

	return aliens
}

func getAlienPowers(e *colly.HTMLElement) []string {
	htmlContent, _ := e.DOM.Html()
	content := removeBrTags(htmlContent)
	powers := strings.Split(content, "<br/>")
	powers = removeEmpty(powers)

	for i, power := range powers {
		if isSurroundedByAnchorTag(power) {
			powers[i] = removeSurroundingAnchorTags(power)
		}
	}

	return powers
}

func removeBrTags(content string) string {
	reg, _ := regexp.Compile("<b>.*</b>")
	return reg.ReplaceAllString(content, "${1}")
}

func removeEmpty(s []string) []string {
	r := []string{}

	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

func isSurroundedByAnchorTag(content string) bool {
	matched, _ := regexp.MatchString("<.*>[a-zA-Z0-9_ ]*<.*>", content)
	return matched
}

func removeSurroundingAnchorTags(content string) string {
	reg, _ := regexp.Compile("<.*>([a-zA-Z0-9_ ]*)<.*>")
	return reg.ReplaceAllString(content, "${1}")
}
