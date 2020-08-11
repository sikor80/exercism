// Package isogram is a solution to exercism.io exercise titled Isogram.
package isogram

import "strings"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(str string) bool {
	set := make(map[rune]bool)
	// Not sure why it doesn't work with the following?
	// var set map[rune]bool

	for _, rune := range strings.ToLower(str) {
		if rune == '-' || rune == ' ' {
			continue
		}
		if _, ok := set[rune]; ok {
			return false
		}
		set[rune] = true
	}
	return true
}
