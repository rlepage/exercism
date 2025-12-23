package scrabble

import (
	"strings"
)

func getLetterValue(letter rune) int {
	switch letter {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

func Score(word string) int {
	var score int = 0
	letters := []rune(strings.ToUpper(word))

	for _, letter := range letters {
		score += getLetterValue(letter)
	}

	return score
}
