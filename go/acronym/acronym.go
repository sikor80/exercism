// Package acronym is a solution to exercism.io exercise titled Acronym.
package acronym

import (
	"log"
	"regexp"
	"strings"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {

	// Remove all non-letters but the apostrophe sign.
	reg, err := regexp.Compile("[^A-Za-z']+")
	if err != nil {
		log.Fatal(err)
	}
	s = reg.ReplaceAllString(s, " ")

	// Get all words.
	words := strings.Fields(s)

	// Get first letter from each word and concat it.
	var result string
	for _, word := range words {
		result += word[0:1]
	}

	// Change to upper case.
	result = strings.ToUpper(result)

	return result
}
