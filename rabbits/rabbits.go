package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RatLad69/rosalind"
)

func main() {
	data, err := rosalind.ReadData("../data/rosalind_fib.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}

	numbers := strings.Fields(data)
	n, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Println("Error getting n:", err)
	}
	k, err := strconv.Atoi(numbers[1])
	if err != nil {
		fmt.Println("Error getting k:", err)
	}
	fmt.Printf("%v generations, k = %v\n", n, k)

	last, current, total := 1, 1, 2
	for i := 2; i < n; i++ {
		total = k*last + current
		last, current = current, total
		fmt.Printf("After %v months, population is %v\n", i+1, total)
	}

	fmt.Println("Total population at end: ", total)
}
