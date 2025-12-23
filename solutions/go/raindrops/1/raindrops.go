package raindrops

import "strconv"

type Suffix struct {
	nb  int
	str string
}

func Convert(number int) string {
	var res string

	suffixes := []Suffix{
		Suffix{nb: 3, str: "Pling"},
		Suffix{nb: 5, str: "Plang"},
		Suffix{nb: 7, str: "Plong"},
	}

	for _, suffix := range suffixes {
		if number%suffix.nb == 0 {
			res += suffix.str
		}
	}

	if res == "" {
		return strconv.Itoa(number)
	}
	return res
}
