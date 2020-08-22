// Package etl is a solution to exercism.io exercise titled ETL.
package etl

import "strings"

// Transform converts Scrabble scores from a legacy system to a new format.
func Transform(given map[int][]string) map[string]int {
	result := make(map[string]int)

	for key, value := range given {
		for _, letter := range value {
			result[strings.ToLower(letter)] = key
		}
	}

	return result
}
