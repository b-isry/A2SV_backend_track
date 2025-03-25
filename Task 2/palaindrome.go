package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	clean := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return -1
		}
		return unicode.ToLower(r)
	}, s)

	for i := range clean {
		if clean[i] != clean[len(clean)-i-1] {
			return false
		}
	}
	return true
}
