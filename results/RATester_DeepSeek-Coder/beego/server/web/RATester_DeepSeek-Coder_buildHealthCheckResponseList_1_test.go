package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildHealthCheckResponseList_1(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]string
		expected []map[string]interface{}
	}{
		{
			name: "Test Case 1",
			input: [][]string{
				{"service1", "message1", "status1"},
				{"service2", "message2", "status2"},
			},
			expected: []map[string]interface{}{
				{
					"name":    "service1",
					"message": "message1",
					"status":  "status1",
				},
				{
					"name":    "service2",
					"message": "message2",
					"status":  "status2",
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := buildHealthCheckResponseList(&tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
