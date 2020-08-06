// Package proverb is a solution to exercism.io exercise titled Proverb.
package proverb

import "fmt"

const (
	line     = "For want of a %s the %s was lost."
	lastLine = "And all for the want of a %s."
)

// Proverb generates slice with a proverb given provided slice of strings.
func Proverb(rhyme []string) []string {
	var proverb []string
	length := len(rhyme)

	switch length {
	case 0:
		proverb = []string{}
	case 1:
		proverb = []string{fmt.Sprintf(lastLine, rhyme[0])}
	default:
		// Thanks to range rhyme[:length-1] we don't use the last word as rhyme[i].
		for i := range rhyme[:length-1] {
			proverb = append(proverb, fmt.Sprintf(line, rhyme[i], rhyme[i+1]))
		}
		proverb = append(proverb, fmt.Sprintf(lastLine, rhyme[0]))
	}

	return proverb
}
