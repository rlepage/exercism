package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	for i := 'a'; i <= 'z'; i++ {
		if !strings.Contains(input, string(i)) {
			return false
		}
	}
	return true
}
