package main

import (
	"fmt"
	"strings"
	"unicode"
)

func LatinLetters(s string) string {
	sb := strings.Builder{}

	for _, r := range s {
		if unicode.Is(unicode.Latin, r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func main() {
	s := "привет world"
	fmt.Println(LatinLetters(s))
}
