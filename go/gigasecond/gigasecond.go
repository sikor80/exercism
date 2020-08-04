// Package gigasecond is a solution to exercism.io exercise titled Gigasecond.
package gigasecond

import "time"

// AddGigasecond adds 1 million seconds to a provided time.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)

	return t
}
