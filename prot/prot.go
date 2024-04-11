package main

import (
	"fmt"

	"github.com/RatLad69/rosalind"
)

func CToS(codon [3]rune) string {
	sCodon := ""
	for _, r := range codon {
		sCodon += string(r)
	}
	return sCodon
}

func main() {
	codons := map[string]rune{
		"UUU": 'F',
		"UUC": 'F',
		"UUA": 'L',
		"UUG": 'L',
		"UCU": 'S',
		"UCC": 'S',
		"UCA": 'S',
		"UCG": 'S',
		"UAU": 'Y',
		"UAC": 'Y',
		"UAA": '@',
		"UAG": '@',
		"UGU": 'C',
		"UGC": 'C',
		"UGA": '@',
		"UGG": 'W',
		"CUU": 'L',
		"CUC": 'L',
		"CUA": 'L',
		"CUG": 'L',
		"CCU": 'P',
		"CCC": 'P',
		"CCA": 'P',
		"CCG": 'P',
		"CAU": 'H',
		"CAC": 'H',
		"CAA": 'Q',
		"CAG": 'Q',
		"CGU": 'R',
		"CGC": 'R',
		"CGA": 'R',
		"CGG": 'R',
		"AUU": 'I',
		"AUC": 'I',
		"AUA": 'I',
		"AUG": 'M',
		"ACU": 'T',
		"ACC": 'T',
		"ACA": 'T',
		"ACG": 'T',
		"AAU": 'N',
		"AAC": 'N',
		"AAA": 'K',
		"AAG": 'K',
		"AGU": 'S',
		"AGC": 'S',
		"AGA": 'R',
		"AGG": 'R',
		"GUU": 'V',
		"GUC": 'V',
		"GUA": 'V',
		"GUG": 'V',
		"GCU": 'A',
		"GCC": 'A',
		"GCA": 'A',
		"GCG": 'A',
		"GAU": 'D',
		"GAC": 'D',
		"GAA": 'E',
		"GAG": 'E',
		"GGU": 'G',
		"GGC": 'G',
		"GGA": 'G',
		"GGG": 'G',
	}

	data, err := rosalind.ReadData("../data/rosalind_prot.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	RNA := rosalind.RmWhiteSpace(data)
	var codon [3]rune
	protein := ""

	for i, base := range RNA {
		codon[i%3] = base
		if i%3 == 2 {
			protein += string(codons[CToS(codon)])
		}
	}

	fmt.Println(protein)
}
