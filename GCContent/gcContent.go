package main

import (
	"fmt"
	"strings"

	"github.com/RatLad69/rosalind"
)

var dnaStrings map[string]string
var label string
var dna string

func getGcContent(dna string) float64 {
	length := 0
	gc := 0

	for _, sym := range dna {
		length++
		if sym == 'G' || sym == 'C' {
			gc++
		}
	}

	return 100 * float64(gc) / float64(length)
}

func main() {
	data, err := rosalind.ReadData("../data/rosalind_gc.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	lines := strings.Fields(data)
	lines = append(lines, ">END")

	dnaStrings = make(map[string]string)
	for i, line := range lines {
		if strings.HasPrefix(line, ">") {
			label = rosalind.RmWhiteSpace(line)
			dna = ""
		} else {
			dna += rosalind.RmWhiteSpace(line)
			if strings.HasPrefix(lines[i+1], ">") {
				dnaStrings[label] = dna
			}
		}
	}

	maxGc, maxString := float64(0), ""
	for dnaName, dnaString := range dnaStrings {
		gc := getGcContent(dnaString)
		if gc > maxGc {
			maxGc = gc
			maxString = dnaName
		}
	}

	fmt.Println(maxString)
	fmt.Println(maxGc)
}
