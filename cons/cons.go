package main

import (
	"fmt"
	"strings"

	"github.com/RatLad69/rosalind"
)

func sliceToString(s []int64, space string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(s), " ", space, -1), "[]")
}

func profToString(profile [][]int64) string {
	pString := ""
	symbols := []string{"A", "C", "G", "T"}
	for i, symbol := range symbols {
		pString += fmt.Sprintf("%s: %s", symbol, sliceToString(profile[i], " "))
		if i != 3 {
			pString += "\n"
		}
	}
	return pString
}

func getConsensus(a, c, g, t int64) (rune, error) {
	if a >= c && a >= g && a >= t {
		return 'A', nil
	}
	if c >= a && c >= g && c >= t {
		return 'C', nil
	}
	if g >= a && g >= c && g >= t {
		return 'G', nil
	}
	if t >= a && t >= c && t >= g {
		return 'T', nil
	}
	return 'X', fmt.Errorf("could not determine consensus")
}

func main() {
	data, err := rosalind.ReadData("../data/rosalind_cons.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	dnaStrings := rosalind.ParseFasta(data)
	firstName := strings.Fields(data)[0]
	dnaLength := len(dnaStrings[firstName])

	stringsM := make([][]rune, len(dnaStrings))
	i := 0
	for _, dnaString := range dnaStrings {
		stringsM[i] = make([]rune, dnaLength)
		for j, c := range dnaString {
			stringsM[i][j] = c
		}
		i++
	}

	profile := make([][]int64, 4)
	for i := 0; i < len(profile); i++ {
		profile[i] = make([]int64, dnaLength)
	}

	var a, c, g, t int64
	consensus := make([]rune, dnaLength)
	for i := 0; i < dnaLength; i++ {
		a, c, g, t = 0, 0, 0, 0
		for j := 0; j < len(stringsM); j++ {
			dnaSym := stringsM[j][i]
			if dnaSym == 'A' {
				a++
			} else if dnaSym == 'C' {
				c++
			} else if dnaSym == 'G' {
				g++
			} else if dnaSym == 'T' {
				t++
			}
		}
		profile[0][i], profile[1][i], profile[2][i], profile[3][i] = a, c, g, t
		c, err := getConsensus(a, c, g, t)
		if err != nil {
			fmt.Println("Error:", err)
		}
		consensus[i] = c
	}
	for _, r := range consensus {
		fmt.Printf("%c", r)
	}
	fmt.Printf("\n")
	fmt.Println(profToString(profile))
}
