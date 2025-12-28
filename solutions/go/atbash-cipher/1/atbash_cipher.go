package atbash

import (
	"regexp"
	"strings"
)

func cypherChar(letter rune) rune {
	if letter >= 'a' && letter <= 'z' {
		letter = 'z' - letter
		if letter < 'a' {
			letter += 'a'
		}
		return letter
	}

	if letter >= '0' && letter <= '9' {
		return letter
	}

	return ' '
}

func Atbash(s string) string {
	reNonAlphanum, _ := regexp.Compile("[^A-Za-z0-9]+")
	s = reNonAlphanum.ReplaceAllString(strings.ToLower(s), "")

	cyphertext := ""
	for i, l := range s {
		if i%5 == 0 && i != 0 {
			cyphertext += " "
		}
		cyphertext += string(cypherChar(l))
	}

	return cyphertext
}
