package wordcount

import (
	"strings"
	"unicode"
)

// This one is pretty good also:
// https://exercism.io/tracks/go/exercises/word-count/solutions/8bd5770114664fb4b3c1c16cee906805

// Frequency represent a count of discovered words.
type Frequency map[string]int

func split(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '\''
}

// WordCount returns the frequency of each word.
func WordCount(phrase string) Frequency {
	words := strings.FieldsFunc(strings.ToLower(phrase), split)
	// f := Frequency{}
	f := make(Frequency)

	for _, w := range words {
		w = strings.Trim(w, "'")
		// f[w] = f[w] + 1
		f[w]++
	}

	return f

}
