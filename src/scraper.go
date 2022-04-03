package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

const (
	ben10FandomDomain1 = "ben10.fandom.com"
	ben10FandomDomain2 = "www.ben10.fandom.com"
	cacheDir           = "./ben10_fandom_cache"
)

const (
	htmlATag            = "<a.*?>.*</a>"
	htmlATagContent     = "<a.*?>(.*)</a>"
	htmlBTag            = "<b>.*</b>"
	htmlBrTag           = "<br/>"
	htmlSmallTag        = "<small>.*</small>"
	htmlSmallTagContent = "<small>(.*)</small>"
)

type alien struct {
	Name       string   `json:"name"`
	ImgURL     string   `json:"imgURL"`
	Species    string   `json:"species"`
	HomePlanet string   `json:"homePlanet"`
	Powers     []string `json:"powers"`
}

func getAliensData(wikiURL string) []alien {
	c := colly.NewCollector(
		colly.AllowedDomains(ben10FandomDomain1, ben10FandomDomain2),
		colly.CacheDir(cacheDir),
	)

	c.OnHTML(".category-page__member-link", func(e *colly.HTMLElement) {
		wikiURL := e.Attr("href")
		e.Request.Visit(wikiURL)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("> Downloading...", r.URL.String())
	})

	aliens := []alien{}

	c.OnHTML(".portable-infobox", func(e *colly.HTMLElement) {
		name := e.ChildText("[data-source=name]")
		imgURL := getAlienImgURL(e)
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

	c.Visit(wikiURL)

	return aliens
}

func getAlienImgURL(e *colly.HTMLElement) string {
	imgURL := e.ChildAttr("img", "src")
	pngIndex := strings.Index(strings.ToLower(imgURL), ".png")
	return imgURL[:pngIndex+4]
}

func getAlienPowers(e *colly.HTMLElement) []string {
	powers := []string{}

	htmlContent, _ := e.DOM.Html()
	content := removeHTMLTag(htmlBTag, htmlContent)

	powers = strings.Split(content, htmlBrTag)
	powers = removeEmpty(powers)

	for i, power := range powers {
		if isSurroundedByHTMLTag(htmlATag, power) {
			powers[i] = removeHTMLTag(htmlATagContent, power)
		}

		if isSurroundedByHTMLTag(htmlSmallTag, power) {
			powers[i] = removeHTMLTag(htmlSmallTagContent, power)
		}

		powers[i] = strings.Trim(powers[i], " ")
	}

	return powers
}

func isSurroundedByHTMLTag(regex, src string) bool {
	matched, _ := regexp.MatchString(regex, src)
	return matched
}

func removeHTMLTag(pattern, src string) string {
	reg := regexp.MustCompile(pattern)
	return reg.ReplaceAllString(src, "${1}")
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
