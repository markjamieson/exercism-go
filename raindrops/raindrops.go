// Package raindrops implements the Convert function
package raindrops

import (
	"fmt"
)

var conversions = []struct {
	factor int
	output string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert returns a string based on the rules outlined in the README.
func Convert(input int) string {
	result := ""
	for _, conversion := range conversions {
		if input%conversion.factor == 0 {
			result = fmt.Sprintf("%v%v", result, conversion.output)
		}
	}
	if result == "" {
		return fmt.Sprintf("%v", input)
	}
	return result
}
