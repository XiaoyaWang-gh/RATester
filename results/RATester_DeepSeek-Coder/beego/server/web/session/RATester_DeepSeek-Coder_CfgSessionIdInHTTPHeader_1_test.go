package session

import (
	"fmt"
	"testing"
)

func TestCfgSessionIdInHTTPHeader_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		enable   bool
		expected bool
	}{
		{
			name:     "Enable",
			enable:   true,
			expected: true,
		},
		{
			name:     "Disable",
			enable:   false,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			config := &ManagerConfig{}
			CfgSessionIdInHTTPHeader(tc.enable)(config)
			if config.EnableSidInHTTPHeader != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, config.EnableSidInHTTPHeader)
			}
		})
	}
}
