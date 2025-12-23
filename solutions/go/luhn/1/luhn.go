package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}

	var sum int
	for idx, char := range id {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}

		if idx%2 != (len(id) % 2) {
			sum += digit
		} else {
			tmp := 2 * digit
			if tmp > 9 {
				tmp -= 9
			}
			sum += tmp
		}
	}

	return sum%10 == 0
}
