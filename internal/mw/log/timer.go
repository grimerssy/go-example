package log

import (
	"time"
)

func startTimer() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}
