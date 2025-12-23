package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var proverb []string = make([]string, len(rhyme))
	var l int = len(rhyme)
	if l == 0 {
		return proverb
	}

	for i := 0; i < l-1; i++ {
		proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	proverb[l-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	return proverb
}
