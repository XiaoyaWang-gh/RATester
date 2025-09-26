package middleware

import (
	"fmt"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestNewRateLimiterMemoryStoreWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	config := RateLimiterMemoryStoreConfig{
		Rate:      10,
		Burst:     10,
		ExpiresIn: 10 * time.Second,
	}

	// WHEN
	store := NewRateLimiterMemoryStoreWithConfig(config)

	// THEN
	assert.Equal(t, 10, store.rate)
	assert.Equal(t, 10, store.burst)
	assert.Equal(t, 10*time.Second, store.expiresIn)
}
