package etl

import "strings"

const NB_LETTERS int8 = 26

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int, NB_LETTERS)

	for score, letters := range in {
		for _, letter := range letters {
			lowercaseLetter := strings.ToLower(letter)
			out[lowercaseLetter] = score
		}
	}

	return out
}
