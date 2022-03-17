package main

import "regexp"

const (
	htmlATag            = "<a.*?>.*</a>"
	htmlATagContent     = "<a.*?>(.*)</a>"
	htmlBTag            = "<b>.*</b>"
	htmlBrTag           = "<br/>"
	htmlSmallTag = "<small>.*</small>"
	htmlSmallTagContent = "<small>(.*)</small>"
)

func isSurroundedByTag(regex, src string) bool {
	matched, err := regexp.MatchString(regex, src)
	if err != nil {
		return false
	}

	return matched
}

func removeHTMLTag(pattern, src string) string {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return ""
	}

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
