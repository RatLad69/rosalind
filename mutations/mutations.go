package main

import (
	"fmt"
	"strings"

	"github.com/RatLad69/rosalind"
)

func main() {
	data, err := rosalind.ReadData("../data/rosalind_hamm.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	dnaStrings := strings.Fields(data)

	for i, s := range dnaStrings {
		dnaStrings[i] = rosalind.RmWhiteSpace(s)
	}

	hamming, err := rosalind.Hamming(dnaStrings[0], dnaStrings[1])
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Hamming distance is", hamming)
}
