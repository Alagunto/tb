package tb

import (
	"errors"
	"time"
)

// RetryPolicy defines the retry behavior for failed API calls.
type RetryPolicy struct {
	// MaxRetries is the maximum number of retry attempts.
	MaxRetries int

	// BackoffFunc calculates the delay before the next retry.
	// It receives the retry attempt number (starting from 1) and the error.
	BackoffFunc func(retry int, err error) time.Duration

	// ShouldRetry determines if an error should trigger a retry.
	// If nil, only FloodError will be retried.
	ShouldRetry func(error) bool
}

// DefaultBackoff implements exponential backoff with jitter.
func DefaultBackoff(retry int, err error) time.Duration {
	// For FloodError, use the RetryAfter value from Telegram
	var floodErr FloodError
	if errors.As(err, &floodErr) {
		return time.Duration(floodErr.RetryAfter) * time.Second
	}

	// Exponential backoff: 2^retry seconds, max 60 seconds
	backoff := time.Duration(1<<uint(retry)) * time.Second
	if backoff > 60*time.Second {
		backoff = 60 * time.Second
	}
	return backoff
}

// DefaultRetryPolicy returns a sensible default retry policy.
func DefaultRetryPolicy() RetryPolicy {
	return RetryPolicy{
		MaxRetries:  3,
		BackoffFunc: DefaultBackoff,
		ShouldRetry: func(err error) bool {
			// Retry on flood errors and 5xx server errors
			var floodErr FloodError
			if errors.As(err, &floodErr) {
				return true
			}

			var apiErr *Error
			if errors.As(err, &apiErr) {
				// Retry on internal server errors
				return apiErr.Code >= 500 && apiErr.Code < 600
			}

			return false
		},
	}
}

// WithRetry wraps a function with retry logic.
func WithRetry[T any](fn func() (T, error), policy RetryPolicy) (T, error) {
	var result T
	var err error

	for attempt := 0; attempt <= policy.MaxRetries; attempt++ {
		result, err = fn()
		if err == nil {
			return result, nil
		}

		// Check if we should retry
		shouldRetry := policy.ShouldRetry != nil && policy.ShouldRetry(err)
		if !shouldRetry || attempt == policy.MaxRetries {
			break
		}

		// Calculate backoff
		delay := policy.BackoffFunc(attempt+1, err)
		time.Sleep(delay)
	}

	return result, err
}

// RateLimiter provides rate limiting for API calls.
type RateLimiter struct {
	// MaxCalls is the maximum number of calls allowed per period.
	MaxCalls int

	// Period is the time window for rate limiting.
	Period time.Duration

	tokens chan struct{}
}

// NewRateLimiter creates a new rate limiter.
// Example: NewRateLimiter(30, time.Second) allows 30 calls per second.
func NewRateLimiter(maxCalls int, period time.Duration) *RateLimiter {
	rl := &RateLimiter{
		MaxCalls: maxCalls,
		Period:   period,
		tokens:   make(chan struct{}, maxCalls),
	}

	// Fill the token bucket
	for i := 0; i < maxCalls; i++ {
		rl.tokens <- struct{}{}
	}

	// Refill tokens periodically
	go func() {
		ticker := time.NewTicker(period / time.Duration(maxCalls))
		defer ticker.Stop()

		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default:
				// Bucket is full
			}
		}
	}()

	return rl
}

// Wait blocks until a token is available.
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

// TryWait attempts to get a token without blocking.
// Returns true if a token was acquired, false otherwise.
func (rl *RateLimiter) TryWait() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}
