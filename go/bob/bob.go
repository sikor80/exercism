// Package bob is a solution to exercism.io exercise named Bob.
package bob

import (
	"strings"
	"unicode"
)

func hasLetters(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func isCapitalized(s string) bool {
	if strings.ToUpper(s) == s {
		return true
	}
	return false
}

func isQuestion(s string) bool {
	if strings.HasSuffix(strings.TrimSpace(s), "?") {
		return true
	}
	return false
}

func isEmpty(s string) bool {
	if strings.TrimSpace(s) == "" {
		return true
	}
	return false
}

// Hey simulates Bob - a lackadaisical teenager. It takes string as input and provides different answers.
func Hey(remark string) string {
	var answer string

	switch {
	// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
	case isCapitalized(remark) && isQuestion(remark) && hasLetters(remark):
		answer = "Calm down, I know what I'm doing!"
	// Bob answers 'Sure.' if you ask him a question, such as "How are you?".
	case isQuestion(remark):
		answer = "Sure."
	// He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
	case isCapitalized(remark) && hasLetters(remark):
		answer = "Whoa, chill out!"
	// He says 'Fine. Be that way!' if you address him without actually saying anything.
	case isEmpty(remark):
		answer = "Fine. Be that way!"
	// He answers 'Whatever.' to anything else.
	default:
		answer = "Whatever."
	}

	return answer
}
