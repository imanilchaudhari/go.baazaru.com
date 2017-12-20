package helpers

import (
	"fmt"
	"strconv"
	"time"
)

// ParseTimeout returns a parsed duration from a string
// A duration string value must be a positive integer, optionally followed by a corresponding time unit (s|m|h).
func ParseTimeout(duration string) (time.Duration, error) {
	if i, err := strconv.ParseInt(duration, 10, 64); err == nil && i >= 0 {
		return (time.Duration(i) * time.Second), nil
	}
	if requestTimeout, err := time.ParseDuration(duration); err == nil {
		return requestTimeout, nil
	}
	return 0, fmt.Errorf("Invalid timeout value. Timeout must be a single integer in seconds, or an integer followed by a corresponding time unit (e.g. 1s | 2m | 3h)")
}
