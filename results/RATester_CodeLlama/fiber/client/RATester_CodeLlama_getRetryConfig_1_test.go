package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestGetRetryConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &core{
		client: &Client{
			retryConfig: &RetryConfig{
				InitialInterval: 1 * time.Second,
			},
		},
	}
	assert.Equal(t, c.getRetryConfig(), &RetryConfig{
		InitialInterval: 1 * time.Second,
	})
}
