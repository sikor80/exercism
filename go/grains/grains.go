// Package grains is a solution to exercism.io exercise titled Grains.
package grains

import (
	"fmt"
	"math"
)

// Square calculates the number of grains of wheat on a chessboard given that the number on each square doubles.
func Square(i int) (uint64, error) {
	if i < 1 || i > 64 {
		return 0, fmt.Errorf("Out of range. Must be between 1 and 64. Got: %v", i)
	}

	// Faster, not importing math package.
	// return 1 << (i - 1), nil

	// More readible.
	return uint64(math.Pow(2, float64(i-1))), nil
}

// Total returns the total number of grains on the 64 squares chessboard.
func Total() uint64 {
	var result uint64

	// Faster, not importing math package.
	// for i := 0; i < 64; i++ {
	// 	sum += 1 << i
	// }

	// More readible.
	for i := 0; i < 64; i++ {
		result += uint64(math.Pow(2, float64(i)))
	}

	return result
}
