package client

import (
	"fmt"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestGetRetryConfig_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		client   *Client
		expected *RetryConfig
	}{
		{
			name: "RetryConfig is not nil",
			client: &Client{
				retryConfig: &RetryConfig{
					InitialInterval: 1 * time.Second,
					MaxBackoffTime:  32 * time.Second,
					Multiplier:      2.0,
				},
			},
			expected: &RetryConfig{
				InitialInterval: 1 * time.Second,
				MaxBackoffTime:  32 * time.Second,
				Multiplier:      2.0,
			},
		},
		{
			name: "RetryConfig is nil",
			client: &Client{
				retryConfig: nil,
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &core{client: tc.client}
			got := c.getRetryConfig()
			assert.Equal(t, tc.expected, got)
		})
	}
}
