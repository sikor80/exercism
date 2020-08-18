// Package luhn is a solution to exercism.io exercise titled Luhn.
package luhn

import (
	"strings"
)

// Valid takes a number to determine whether or not it is valid per the Luhn formula.
func Valid(num string) bool {
	// Spaces are allowed in the input, but they should be stripped before checking.
	num = strings.ReplaceAll(num, " ", "")

	// Strings of length 1 or less are not valid.
	if len(num) < 2 {
		return false
	}

	evenLength := len(num)%2 == 0
	sum := 0

	for _, r := range num {
		// All other non-digit characters are disallowed.
		// Convert rune to int.
		digit := int(r - '0')
		if digit < 0 || digit > 9 {
			return false
		}

		if evenLength {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}

		sum += digit
		evenLength = !evenLength
	}

	return sum%10 == 0
}
