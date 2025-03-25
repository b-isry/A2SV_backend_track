package main

import (
	"strings"
	"unicode"
)

func frequency(s string) map[string]int {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1 
		}
		return unicode.ToLower(r)
	}, s)

	arr := strings.Fields(clean)
	word := make(map[string]int)

	for i := range arr {
		word[arr[i]]++
	}
	return word
}
