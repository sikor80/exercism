// Package strand is a solution to exercism.io exercise titled RNA Transcription.
package strand

// ToRNA converts a DNA strand, to its RNA complement (per RNA transcription).
func ToRNA(dna string) string {
	// Convert to slice of runes (or sometimes bytes) when working with strings.
	dnaRunes := []rune(dna)

	// Strings are immutable so create a new one and populate it.
	rnaRunes := make([]rune, len(dnaRunes))

	for i, r := range dnaRunes {
		switch r {
		case 'G':
			rnaRunes[i] = 'C'
		case 'C':
			rnaRunes[i] = 'G'
		case 'T':
			rnaRunes[i] = 'A'
		case 'A':
			rnaRunes[i] = 'U'
		default:
			rnaRunes[i] = r
		}
	}
	return string(rnaRunes)

}
