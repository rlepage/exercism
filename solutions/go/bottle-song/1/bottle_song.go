package bottlesong

import (
	"fmt"
	"strings"
)

var Template []string = []string{
	"%s hanging on the wall,",
	"%s hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be %s hanging on the wall.",
}

var NbOfBottles []string

func InitNbOfBottles() {
	NbOfBottles = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	for idx, val := range NbOfBottles {
		NbOfBottles[idx] = val + " green bottles"
	}
	NbOfBottles[1] = "one green bottle"
}

func Recite(startBottles, takeDown int) []string {
	var song []string
	InitNbOfBottles()

	for i := startBottles; i > startBottles-takeDown; i-- {
		currentNumber := strings.ToUpper(NbOfBottles[i][:1]) + NbOfBottles[i][1:]

		song = append(song,
			fmt.Sprintf(Template[0], currentNumber),
			fmt.Sprintf(Template[1], currentNumber),
			Template[2],
			fmt.Sprintf(Template[3], NbOfBottles[i-1]),
			"",
		)
	}

	return song[:len(song)-1]
}
