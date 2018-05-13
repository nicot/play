package main

import "testing"

func sort(s string) string {
	return s
}

func uniqChars(s string) bool {
	seenChars := make(map[rune]bool)

	for _, c := range s {
		if seenChars[c] {
			return false
		}
		seenChars[c] = true
	}
	return true
}

func uniqCharsNoExtra(s string) bool {
	sorted := sort(s)
	var prev rune

	for _, c := range sorted {
		if c == prev {
			return false
		}
	}
	return true
}

func reverse(s []rune) []rune {
	for i := 0; i < len(s)/2; i++ {
		endIndex := len(s) - 1 - i
		tmp := s[endIndex]
		s[endIndex] = s[i]
		s[i] = tmp
	}
	return s
}

func TestReverse(t *testing.T) {
	s := []rune("abcef")
	s = reverse(s)
	t.Error(s)
}

func charCount(s string) map[rune]int {

}

func isPermutation(a, b string) bool {
	ac := charCount(a)
	bc := charCount(b)
	for char, count := range ac {
		// if
	}
}
