package main

import (
	"fmt"
	"strings"
)

func main() {

	var length, delta int
	var input string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	sb := strings.Builder{}
	for _, r := range input {
		sb.WriteRune(nextRune(r, delta))
	}
	fmt.Println(sb.String())
}

func nextRune(r rune, delta int) rune {
	var ret rune
	if r >= 'a' && r <= 'z' {
		ret = rotate(r, delta, 'a')
	} else if r >= 'A' && r <= 'Z' {
		ret = rotate(r, delta, 'A')
	} else {
		ret = r
	}

	return ret
}

const alphabetRange = 'a' - ('z' + rune(1))

func rotate(r rune, delta int, base rune) rune {
	return ((r - base + rune(delta)) % alphabetRange) + base
}
