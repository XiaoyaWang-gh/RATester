package metrics

import (
	"fmt"
	"testing"
)

func TestHasService_1(t *testing.T) {
	testCases := []struct {
		name        string
		serviceName string
		expected    bool
	}{
		{
			name:        "Test Case 1: Service Exists",
			serviceName: "service1",
			expected:    true,
		},
		{
			name:        "Test Case 2: Service Does Not Exist",
			serviceName: "service2",
			expected:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &dynamicConfig{
				services: map[string]map[string]bool{
					"service1": {},
				},
			}

			result := d.hasService(tc.serviceName)

			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
