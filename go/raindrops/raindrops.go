// Package raindrops is a solution to exercism.io exercise titled Raindrops.
package raindrops

import (
	"bytes"
	"strconv"
)

const (
	factor3 = "Pling"
	factor5 = "Plang"
	factor7 = "Plong"
)

// Convert converts number into a string that contains raindrop sounds corresponding to certain potential factors.
// A factor is a number that evenly divides into another number, leaving no remainder.
func Convert(number int) string {
	var result bytes.Buffer

	if number%3 == 0 {
		result.WriteString(factor3)
	}
	if number%5 == 0 {
		result.WriteString(factor5)
	}
	if number%7 == 0 {
		result.WriteString(factor7)
	}
	if result.Len() == 0 {
		result.WriteString(strconv.Itoa(number))
	}
	return result.String()

}
