package rosalind

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func RmWhiteSpace(input string) string {
	var b strings.Builder
	b.Grow(len(input))
	for _, ch := range input {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func ReadData(dataFile string) (string, error) {
	dataset, err := os.ReadFile(dataFile)
	if err != nil {
		return "", fmt.Errorf("could not read file")
	}

	return string(dataset), nil
}

func Hamming(s1, s2 string) (int64, error) {
	if len(s1) != len(s2) {
		return 0, fmt.Errorf("strings not equal length")
	}

	diffs := int64(0)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffs++
		}
	}

	return diffs, nil
}

func ParseFasta(raw string) map[string]string {
	lines := strings.Fields(raw)
	lines = append(lines, ">END")

	var label, dna string
	dnaStrings := make(map[string]string)
	for i, line := range lines {
		if strings.HasPrefix(line, ">") {
			label = RmWhiteSpace(line)
			dna = ""
		} else {
			dna += RmWhiteSpace(line)
			if strings.HasPrefix(lines[i+1], ">") {
				dnaStrings[label] = dna
			}
		}
	}

	return dnaStrings
}
