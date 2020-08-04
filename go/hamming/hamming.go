// Package hamming is a solution to exercism.io exercise titled Hamming.
package hamming

import "errors"

// Distance calculates the Hamming Distance from two strings.
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("strings must have equal length")
	}

	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
