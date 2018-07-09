// Package acronym implements an abbreviate function for generating acronyms
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns an acronym representation of the input string.
func Abbreviate(s string) string {
	result := ""
	words := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	for i := range words {
		result = strings.Join([]string{result, strings.ToUpper(string(words[i][0]))}, "")
	}
	return result
}
