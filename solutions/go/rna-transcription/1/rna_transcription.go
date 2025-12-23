package strand

var rnaComplement map[rune]rune = map[rune]rune{
	'A': 'U',
	'C': 'G',
	'G': 'C',
	'T': 'A',
}

func ToRNA(dna string) string {
	dnaRunes := []rune(dna)
	for i, nucl := range dnaRunes {
		dnaRunes[i] = rnaComplement[nucl]
	}
	return string(dnaRunes)
}
