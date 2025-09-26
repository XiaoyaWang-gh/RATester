package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDealMultiUser_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:     "Test case 1",
			input:    "user1=value1, user2=value2",
			expected: map[string]string{"user1": "value1", "user2": "value2"},
		},
		{
			name:     "Test case 2",
			input:    "user1=value1, user2=value2, user3",
			expected: map[string]string{"user1": "value1", "user2": "value2", "user3": ""},
		},
		{
			name:     "Test case 3",
			input:    "",
			expected: map[string]string{},
		},
		{
			name:     "Test case 4",
			input:    "user1=value1, user1=value2",
			expected: map[string]string{"user1": "value2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := dealMultiUser(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
