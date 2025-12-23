package isogram

import "strings"

func IsIsogram(word string) bool {
	lettersExist := map[rune]bool{}
	word = strings.ToLower(word)
	chars := []rune(word)
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c != '-' && c != ' ' {
			_, exists := lettersExist[c]
			if exists {
				return false
			}
			lettersExist[c] = true
		}
	}
	return true
}
