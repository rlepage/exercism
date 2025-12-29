package lsproduct

import (
	"errors"
	"regexp"
	"strconv"
)

func rune2int(r rune) int {
	i, _ := strconv.ParseInt(string(r), 10, 64)
	return int(i)
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span <= 0 {
		return 0, errors.New("span must not be negative")
	}
	if span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	}
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(digits) {
		return 0, errors.New("digits input must only contain digits")
	}

	maxProduct := 0
	for i := 0; i <= len(digits)-span; i++ {
		product := 1
		// fmt.Printf("Digits %s span %d - %d %d - %s \n", digits, span, i, i+span, digits[i:i+span])
		for _, digit := range digits[i : i+span] {
			product *= rune2int(rune(digit))
		}
		if product > maxProduct {
			maxProduct = product
		}
	}

	return int64(maxProduct), nil
}
