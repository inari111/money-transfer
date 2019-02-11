package domain

import "time"

type CurrentTimeFunc func() time.Time

// UTC
func NewCurrentTimeFunc() CurrentTimeFunc {
	return time.Now
}
