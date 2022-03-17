package main

import "fmt"

type alien struct {
	name       string
	imgURL     string
	species    string
	homePlanet string
	powers     []string
}

func main() {
	classicAliens := getAliens("https://ben10.fandom.com/wiki/Category:Original_Series_Aliens")
	fmt.Println(classicAliens)
}
