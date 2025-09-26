package common

import (
	"fmt"
	"testing"
)

func TestGetExternalIp_1(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
	}{
		{"Test case 1", "127.0.0.1"},
		{"Test case 2", "192.168.1.1"},
		{"Test case 3", "8.8.8.8"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			externalIp = tc.expected
			result := GetExternalIp()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
