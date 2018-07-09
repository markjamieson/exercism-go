// Package bob defines the Hey function
package bob

import (
	"regexp"
	"strings"
)

// Hey takes a remark as input and returns 1 of 5 responses as defined in the README
func Hey(remark string) string {
	var response string
	var hasCaps = regexp.MustCompile(`[A-Z]+`)
	remark = strings.TrimSpace(remark)
	allCaps := hasCaps.MatchString(remark) && strings.ToUpper(remark) == remark
	isQuestion := remark != "" && string(remark[len(remark)-1]) == "?"
	if remark == "" {
		response = "Fine. Be that way!"
	} else if allCaps && isQuestion {
		response = "Calm down, I know what I'm doing!"
	} else if allCaps && !isQuestion {
		response = "Whoa, chill out!"
	} else if !allCaps && isQuestion {
		response = "Sure."
	} else {
		response = "Whatever."
	}
	return response
}
