package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	reWords := regexp.MustCompile("([a-z]+'[a-z]+|[a-z]+|[0-9]+)")
	words := reWords.FindAllString(strings.ToLower(phrase), -1)

	var frequencies Frequency = map[string]int{}
	for _, word := range words {
		frequencies[word]++
	}
	return frequencies
}
