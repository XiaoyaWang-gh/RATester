package middleware

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
	"golang.org/x/time/rate"
)

func TestNewRateLimiterMemoryStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rate := rate.Limit(10)

	// when
	store := NewRateLimiterMemoryStore(rate)

	// then
	assert.NotNil(t, store)
	assert.Equal(t, rate, store.rate)
}
