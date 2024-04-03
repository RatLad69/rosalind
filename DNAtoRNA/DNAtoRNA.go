package main

import (
	"fmt"
	"strings"

	"github.com/RatLad69/rosalind"
)

func main() {
	data, err := rosalind.ReadData("../data/rosalind_rna.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	var RNAbuilder strings.Builder
	for _, ch := range data {
		if ch == 'T' {
			RNAbuilder.WriteRune('U')
		} else {
			RNAbuilder.WriteRune(ch)
		}
	}
	fmt.Println(RNAbuilder.String())
}
