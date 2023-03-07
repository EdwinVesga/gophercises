package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer := 1
	for _, r := range input {
		if isUpper(r) {
			answer++
		}
	}
	fmt.Printf("Number of Words: %d\n", answer)
}

func isUpper(r rune) bool {
	min, max := 'A', 'Z'
	var ret bool
	if r >= min && r <= max {
		ret = true
	}

	return ret
}
