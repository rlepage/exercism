package protein

import (
	"errors"
	"fmt"
)

var ErrStop error = errors.New("ErrStop")
var ErrInvalidBase error = errors.New("ErrInvalidBase")

var proteins map[string]string = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	nbCodons := len(rna) / 3
	res := []string{}

	for i := 0; i < nbCodons; i++ {
		codon := fmt.Sprintf("%c%c%c", rna[i*3], rna[i*3+1], rna[i*3+2])
		protein, err := FromCodon(codon)
		if err == ErrInvalidBase {
			return nil, ErrInvalidBase
		}
		if err == ErrStop {
			return res, nil
		}

		res = append(res, protein)
	}

	return res, nil
}

func FromCodon(codon string) (string, error) {
	protein, exists := proteins[codon]
	if !exists {
		return "", ErrInvalidBase
	}
	if proteins[codon] == "STOP" {
		return "", ErrStop
	}

	return protein, nil
}
