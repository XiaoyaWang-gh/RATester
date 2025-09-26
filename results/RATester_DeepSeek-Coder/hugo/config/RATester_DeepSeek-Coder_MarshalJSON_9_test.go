package config

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_9(t *testing.T) {
	type testStruct struct {
		Field string `json:"field"`
	}

	testCases := []struct {
		name     string
		input    ConfigNamespace[testStruct, testStruct]
		expected string
	}{
		{
			name: "Test case 1",
			input: ConfigNamespace[testStruct, testStruct]{
				SourceStructure: testStruct{Field: "test"},
				SourceHash:      "hash",
			},
			expected: `{"field":"test"}`,
		},
		{
			name: "Test case 2",
			input: ConfigNamespace[testStruct, testStruct]{
				SourceStructure: testStruct{Field: "another test"},
				SourceHash:      "another hash",
			},
			expected: `{"field":"another test"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := tc.input.MarshalJSON()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if string(result) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, string(result))
			}
		})
	}
}
