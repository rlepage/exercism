package rotationalcipher

import "strings"

func rotateLetter(letter rune, shiftKey int) rune {
	if letter >= 'a' && letter <= 'z' {
		return ((letter - 'a' + rune(shiftKey)) % 26) + 'a'
	} else if letter >= 'A' && letter <= 'Z' {
		return ((letter - 'A' + rune(shiftKey)) % 26) + 'A'
	}
	return letter
}

func RotationalCipher(plain string, shiftKey int) string {
	str := []string{}
	for _, r := range plain {
		str = append(str, string(rotateLetter(r, shiftKey)))
	}

	return strings.Join(str, "")
}
