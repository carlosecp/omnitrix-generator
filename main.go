package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type alien struct {
	name       string
	imgURL     string
	species    string
	homePlanet string
	powers     []string
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("ben10.fandom.com", "www.ben10.fandom.com"),
		colly.CacheDir("./ben10_fandom_cache"),
	)

	c.OnHTML(".category-page__member-link", func(e *colly.HTMLElement) {
		wikiURL := e.Attr("href")
		if strings.Index(wikiURL, "Ghostfreak") > -1 {
			e.Request.Visit(wikiURL)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	c.OnHTML(".portable-infobox", func(e *colly.HTMLElement) {
		name := e.ChildText("[data-source=name]")
		imgURL := e.ChildAttr("img", "src")
		species := e.ChildText("[data-source=species] a")
		homePlanet := e.ChildText("[data-source=species] a")

		powers := []string{}

		// fmt.Println(e.ChildText("[data-source=power] div").Html())
		e.ForEach("[data-source=power] div", func(_ int, el *colly.HTMLElement) {
			regRemoveBTag, _ := regexp.Compile("<b>.*</b>")
			regRemoveATag, _ := regexp.Compile("<.*>([a-zA-Z0-9_ ]*)<.*>")

			rawContent, _ := el.DOM.Html()
			rawPowers := regRemoveBTag.ReplaceAllString(rawContent, "${1}")
			powers1 := strings.Split(rawPowers, "<br/>")

			for _, power := range powers1 {
				if matched, _ := regexp.MatchString("<.*>[a-zA-Z0-9_ ]*<.*>", power); matched {
					power = regRemoveATag.ReplaceAllString(power, "${1}")
				}

				powers = append(powers, power)
			}
		})

		newAlien := alien{
			name:       name,
			imgURL:     imgURL,
			species:    species,
			homePlanet: homePlanet,
			powers:     powers,
		}

		fmt.Println(newAlien)
	})

	c.Visit("https://ben10.fandom.com/wiki/Category:Original_Series_Aliens")
}

func getAlienPowers(e *colly.HTMLElement) []string {
	content, _ := el.DOM.Html()
	content = sanitizeAlientPowersContent(content)
	fmt.Println(content)
}

func sanitizeAlientPowersContent(content string) string {
	content = removeBrTags(content)
	content = removeAnchorTags(content)
	return content
}

func removeBrTags(content string) string {
	reg, _ := regexp.Compile("<b>.*</b>")
	return reg.ReplaceAllString(content, "${1}")
}

func removeAnchorTags(content string) string {
	reg
}
