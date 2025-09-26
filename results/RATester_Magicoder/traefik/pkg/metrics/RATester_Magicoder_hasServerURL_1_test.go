package metrics

import (
	"fmt"
	"testing"
)

func TestHasServerURL_1(t *testing.T) {
	tests := []struct {
		name          string
		serviceName   string
		serverURL     string
		expected      bool
		dynamicConfig *dynamicConfig
	}{
		{
			name:        "Test Case 1",
			serviceName: "service1",
			serverURL:   "http://server1.com",
			expected:    true,
			dynamicConfig: &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {"http://server1.com": true},
				},
			},
		},
		{
			name:        "Test Case 2",
			serviceName: "service2",
			serverURL:   "http://server2.com",
			expected:    false,
			dynamicConfig: &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {"http://server1.com": true},
				},
			},
		},
		{
			name:        "Test Case 3",
			serviceName: "service1",
			serverURL:   "http://server2.com",
			expected:    false,
			dynamicConfig: &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {"http://server1.com": true},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.dynamicConfig.hasServerURL(test.serviceName, test.serverURL)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
