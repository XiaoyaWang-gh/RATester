package failover

import (
	"context"
	"fmt"
	"testing"
)

func TestSetFallbackHandlerStatus_1(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name     string
		up       bool
		expected bool
	}{
		{
			name:     "up",
			up:       true,
			expected: true,
		},
		{
			name:     "down",
			up:       false,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &Failover{}
			f.SetFallbackHandlerStatus(ctx, test.up)

			if f.fallbackStatus != test.expected {
				t.Errorf("Expected fallbackStatus to be %v, but got %v", test.expected, f.fallbackStatus)
			}
		})
	}
}
