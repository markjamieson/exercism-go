// Package hamming implements Distance function
package hamming

import (
	"errors"
)

// Distance returns the hamming distance between two strings.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Input strings must be same length")
	}
	result := 0
	for i := range a {
		if a[i] != b[i] {
			result += 1
		}
	}
	return result, nil
}
