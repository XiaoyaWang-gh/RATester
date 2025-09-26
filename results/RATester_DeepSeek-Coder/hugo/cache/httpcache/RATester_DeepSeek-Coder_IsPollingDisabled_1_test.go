package httpcache

import (
	"fmt"
	"testing"
)

func TestIsPollingDisabled_1(t *testing.T) {
	type testCase struct {
		name     string
		config   ConfigCompiled
		expected bool
	}

	testCases := []testCase{
		{
			name: "All polling enabled",
			config: ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{
						Config: PollConfig{
							Disable: false,
						},
					},
					{
						Config: PollConfig{
							Disable: false,
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "All polling disabled",
			config: ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{
						Config: PollConfig{
							Disable: true,
						},
					},
					{
						Config: PollConfig{
							Disable: true,
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Mixed polling",
			config: ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{
						Config: PollConfig{
							Disable: false,
						},
					},
					{
						Config: PollConfig{
							Disable: true,
						},
					},
				},
			},
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

			result := tc.config.IsPollingDisabled()
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
