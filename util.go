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
