package main

import (
	"strings"
	"unicode"
)

func FreqCount(line string) map[string]int {
	curr := []rune{}
	m := make(map[string]int)

	for _, r := range line {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {

			curr = append(curr, r)
		} else {
			if len(curr) != 0 {
				m[strings.ToLower(string(curr))]++
				curr = []rune{}
			}
		}
	}
	if len(curr) != 0 {
		m[strings.ToLower(string(curr))]++
		curr = []rune{}
	}
	return m
}

func PalindromeCheck(s string) (ret bool) {
	stripped := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			stripped = append(stripped, unicode.ToLower(r))
		}
	}
	n := len(stripped)

	for i := 0; i < n/2; i++ {
		if stripped[i] != stripped[n-i-1] {
			return false
		}
	}

	return true
}

func main() {
}
