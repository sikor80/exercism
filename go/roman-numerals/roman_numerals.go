// Package romannumerals is a solution to exercism.io exercise titled Roman Numerals.
package romannumerals

import (
	"errors"
)

var (
	base   = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// ToRomanNumeral converts integers to Roman Numerals.
func ToRomanNumeral(number int) (result string, err error) {
	if number <= 0 || number > 3000 {
		return "", errors.New("Invalid input: has to be larger than 0 or equal to or smaller than 3000")
	}
	// Not mine. Taken from https://exercism.io/tracks/go/exercises/roman-numerals/solutions/63fa444d3bae465a9409991943723fdd.
	for i, b := range base {
		q := number / b
		number %= b

		for j := 0; j < q; j++ {
			result += symbol[i]

		}
	}
	return result, nil
}
