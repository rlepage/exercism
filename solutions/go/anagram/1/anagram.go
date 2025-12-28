package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	s := []rune(strings.ToLower(subject))
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	anagrams := []string{}
	for _, candidate := range candidates {
		c := []rune(strings.ToLower(candidate))
		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})

		if string(s) == string(c) && !strings.EqualFold(candidate, subject) {
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}
