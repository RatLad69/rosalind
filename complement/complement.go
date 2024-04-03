package main

import (
	"fmt"

	"github.com/RatLad69/rosalind"
)

func complement(r rune) (rune, error) {
	switch r {
	case 'A':
		return 'T', nil
	case 'T':
		return 'A', nil
	case 'C':
		return 'G', nil
	case 'G':
		return 'C', nil
	default:
		return 'X', fmt.Errorf("unexpected letter in DNA string %q", r)
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		first, err1 := complement(runes[i])
		last, err2 := complement(runes[j])
		if err1 != nil || err2 != nil {
			fmt.Println("Error:", err1, err2)
		}
		runes[i], runes[j] = last, first
	}

	return string(runes)
}

func main() {
	data, err := rosalind.ReadData("../data/rosalind_revc.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	data = rosalind.RmWhiteSpace(data)

	fmt.Println(reverseString(data))
}
