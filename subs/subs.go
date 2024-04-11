package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RatLad69/rosalind"
)

func main() {
	data, err := rosalind.ReadData("../data/rosalind_subs.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	dnaStrings := strings.Fields(data)
	for i := 0; i < len(dnaStrings); i++ {
		dnaStrings[i] = rosalind.RmWhiteSpace(dnaStrings[i])
	}

	foundAt := ""
	for i := 0; i < len(dnaStrings[0])-len(dnaStrings[1])+1; i++ {
		// fmt.Printf("%c", dnaStrings[0][i])
		for j := 0; j < len(dnaStrings[1]); j++ {
			if dnaStrings[0][i+j] != dnaStrings[1][j] {
				break
			} else {
				if j == len(dnaStrings[1])-1 {
					foundAt += strconv.Itoa(i+1) + " "
				}
			}
		}
	}

	fmt.Println(foundAt)
}
