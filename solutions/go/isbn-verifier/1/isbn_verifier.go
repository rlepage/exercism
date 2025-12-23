package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	isbnRunes := []rune(isbn)
	if len(isbnRunes) != 10 {
		return false
	}
	if isbnRunes[9] == 'X' {
		isbnRunes[9] = '9' + 1
	}

	var n int
	for i, r := range isbnRunes {
		n += (10 - i) * int(r-'0')
	}

	return n%11 == 0
}
