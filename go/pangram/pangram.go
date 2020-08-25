package pangram

import "unicode"

const alphabetSize = 26

// IsPangram tests if a sentence is using every letter of the alphabet at least once.
func IsPangram(text string) bool {
	lettersMap := make(map[rune]bool)

	for _, r := range text {
		if unicode.IsLetter(r) {
			lettersMap[unicode.ToLower(r)] = true
		}
	}

	return len(lettersMap) == alphabetSize
}
