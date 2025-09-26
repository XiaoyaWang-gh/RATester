package middleware

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestAllow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	store := &RateLimiterMemoryStore{
		visitors:    map[string]*Visitor{},
		rate:        1,
		burst:       1,
		expiresIn:   1 * time.Minute,
		lastCleanup: time.Now(),
		timeNow: func() time.Time {
			return time.Now()
		},
	}
	identifier := "test"

	// when
	allowed, err := store.Allow(identifier)

	// then
	assert.NoError(t, err)
	assert.True(t, allowed)
}
