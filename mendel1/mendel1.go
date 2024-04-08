package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RatLad69/rosalind"
)

func main() {
	data, err := rosalind.ReadData("../data/rosalind_iprb.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	population := strings.Fields(data)
	var kmn [3]int64
	for i := 0; i < len(population); i++ {
		intz, _ := strconv.Atoi(population[i])
		kmn[i] = int64(intz)
	}

	k, m, n := kmn[0], kmn[1], kmn[2]
	totalPop := k + m + n

	pNN := float64(n*(n-1)) / float64(totalPop*(totalPop-1))
	pMM := float64(m*(m-1)) / float64(totalPop*(totalPop-1))
	pMN := float64(2*m*n) / float64(totalPop*(totalPop-1))
	p := 1 - pNN - .25*pMM - .5*pMN
	fmt.Println("P =", p)
}
