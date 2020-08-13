// Package isogram is a solution to exercism.io exercise titled Isogram.
package isogram

import "unicode"

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(str string) bool {
	set := make(map[rune]bool)

	for _, r := range str {
		if r == '-' || r == ' ' {
			continue
		}
		if set[unicode.ToLower(r)] {
			return false
		}
		set[unicode.ToLower(r)] = true
	}
	return true
}
