package reverse

import "strings"

func Reverse(input string) string {
	var runes = []rune(input)
	var reversed strings.Builder

	for i := len(runes) - 1; i >= 0; i-- {
		reversed.WriteRune(runes[i])
	}

	return reversed.String()
}
