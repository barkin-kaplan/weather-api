package tickerhelper

import (
	"context"
	"time"
)

// RunTickerWithImmediate executes `fn` immediately, then on every `interval`.
func RunTickerWithImmediate(interval time.Duration, fn func()) {
	// Run immediately
	fn()

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		fn()
	}
}

func RunPreciseTicker(ctx context.Context, period time.Duration, task func()) {
	next := time.Now().Truncate(period).Add(period)
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Until(next)):
			task()
			next = next.Add(period)
		}
	}
}
