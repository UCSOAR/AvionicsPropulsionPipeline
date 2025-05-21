package utils

import "time"

func DurationToSeconds(duration time.Duration) int {
	seconds := int(duration.Seconds())
	return seconds
}
