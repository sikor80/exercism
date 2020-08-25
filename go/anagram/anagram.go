package anagram

import (
	"sort"
	"strings"
)

func sortLetters(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

// Detect returns list of anagrams given a word and a list of candidates.
// An anagram is a rearrangement of letters to form a new word.
func Detect(word string, candidates []string) []string {
	var result []string
	loweredWord := strings.ToLower(word)
	sortedWord := sortLetters(loweredWord)

	for _, candidate := range candidates {
		loweredCandidate := strings.ToLower(candidate)
		sortedCandidate := sortLetters(loweredCandidate)

		//  Words are not anagrams of themselves.
		if (sortedCandidate == sortedWord) && (loweredCandidate != loweredWord) {
			result = append(result, candidate)
		}
	}

	return result
}
