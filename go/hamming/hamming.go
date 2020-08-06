// Package hamming is a solution to exercism.io exercise titled Hamming.
package hamming

import "errors"

// Distance calculates the Hamming Distance from two strings.
func Distance(a, b string) (distance int, err error) {
	aRune := []rune(a)
	bRune := []rune(b)

	if len(aRune) != len(bRune) {
		return 0, errors.New("strings must have equal length")
	}

	for i := range aRune {
		if aRune[i] != bRune[i] {
			distance++
		}
	}

	return distance, nil
}
