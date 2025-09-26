package session

import (
	"fmt"
	"testing"
	"time"
)

func TestSetExpiry_1(t *testing.T) {
	s := &Session{
		exp: 1 * time.Hour,
	}

	testCases := []struct {
		name     string
		input    time.Duration
		expected time.Duration
	}{
		{
			name:     "SetExpiry with 2 hours",
			input:    2 * time.Hour,
			expected: 2 * time.Hour,
		},
		{
			name:     "SetExpiry with 3 hours",
			input:    3 * time.Hour,
			expected: 3 * time.Hour,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s.SetExpiry(tc.input)
			if s.exp != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, s.exp)
			}
		})
	}
}
