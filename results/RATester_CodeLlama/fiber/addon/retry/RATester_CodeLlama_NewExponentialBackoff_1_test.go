package retry

import (
	"fmt"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestNewExponentialBackoff_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	config := []Config{
		{
			InitialInterval: 1 * time.Second,
			MaxBackoffTime:  32 * time.Second,
			Multiplier:      2.0,
			MaxRetryCount:   10,
		},
	}

	// when
	eb := NewExponentialBackoff(config...)

	// then
	assert.Equal(t, 1*time.Second, eb.InitialInterval)
	assert.Equal(t, 32*time.Second, eb.MaxBackoffTime)
	assert.Equal(t, 2.0, eb.Multiplier)
	assert.Equal(t, 10, eb.MaxRetryCount)
	assert.Equal(t, 1*time.Second, eb.currentInterval)
}
