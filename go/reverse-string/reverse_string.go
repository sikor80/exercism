package reverse

import (
	"strings"
	"unicode/utf8"
)

// Reverse take a string and returns it reversed.
func Reverse(text string) string {
	size := utf8.RuneCountInString(text)
	runes := []rune(text)
	var reversed strings.Builder

	for i := size - 1; i >= 0; i-- {
		reversed.WriteRune(runes[i])
	}
	return reversed.String()

}
