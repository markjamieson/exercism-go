// Package gigasecond has one function for adding Gigaseconds to Times
package gigasecond

import "time"
import "math"

// AddGigasecond takes a Time and returns a new Time 10^9 seconds in the future
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Duration(math.Pow10(9)) * time.Second)
}
