package bob

import "strings"

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	var upper, lower, digits int
	for _, r := range []rune(remark) {
		if r >= 'a' && r <= 'z' {
			lower++
		} else if r >= 'A' && r <= 'Z' {
			upper++
		} else if r >= '0' && r <= '9' {
			digits++
		}
	}

	var isQuestion bool = strings.HasSuffix(remark, "?")
	var isScreaming bool = upper > lower
	var isEmpty bool = upper == 0 && lower == 0 && digits == 0

	switch {
	case isQuestion && isScreaming:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isScreaming:
		return "Whoa, chill out!"
	case isEmpty:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
