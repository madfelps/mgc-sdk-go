package retry

import (
	"math"
	"time"
)

func ShouldRetry(statusCode int) bool {
	return statusCode >= 500 || statusCode == 429
}

func GetNextBackoff(attempt int, backoffFactor float64, initialInterval, maxInterval time.Duration) time.Duration {
	multiplier := math.Pow(backoffFactor, float64(attempt))
	backoffDuration := initialInterval * time.Duration(multiplier)

	if backoffDuration > maxInterval {
		backoffDuration = maxInterval
	}

	return backoffDuration
}
