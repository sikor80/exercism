// Package scrabble is a solution to exercism.io exercise titled Scrabble Score.
package scrabble

import "strings"

// Score computes the Scrabble score for a provided word.
func Score(str string) int {
	var result int
	for _, r := range str {
		switch strings.ToUpper(string(r)) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			result = result + 1
		case "D", "G":
			result = result + 2
		case "B", "C", "M", "P":
			result = result + 3
		case "F", "H", "V", "W", "Y":
			result = result + 4
		case "K":
			result = result + 5
		case "J", "X":
			result = result + 8
		case "Q", "Z":
			result = result + 10
		default:
			result = result + 0
		}
	}

	return result
}
