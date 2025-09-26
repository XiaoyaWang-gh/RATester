package docker

import (
	"fmt"
	"testing"
)

func TestGetServiceName_2(t *testing.T) {
	tests := []struct {
		name     string
		input    dockerData
		expected string
	}{
		{
			name: "Test case 1",
			input: dockerData{
				ServiceName: "service1",
				Labels: map[string]string{
					labelDockerComposeProject: "project1",
					labelDockerComposeService: "service2",
				},
			},
			expected: "service2_project1",
		},
		{
			name: "Test case 2",
			input: dockerData{
				ServiceName: "service1",
				Labels: map[string]string{
					"randomLabel": "randomValue",
				},
			},
			expected: "service1",
		},
		{
			name: "Test case 3",
			input: dockerData{
				ServiceName: "service1",
				Labels:      map[string]string{},
			},
			expected: "service1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := getServiceName(tt.input)
			if got != tt.expected {
				t.Errorf("getServiceName() = %v, want %v", got, tt.expected)
			}
		})
	}
}
