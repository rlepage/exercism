package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("invalid input")
	}

	var roman string
	type Digit struct {
		value  int
		letter string
	}
	digits := []Digit{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, digit := range digits {
		for input >= digit.value {
			roman += digit.letter
			input -= digit.value
		}
	}

	return roman, nil
}
