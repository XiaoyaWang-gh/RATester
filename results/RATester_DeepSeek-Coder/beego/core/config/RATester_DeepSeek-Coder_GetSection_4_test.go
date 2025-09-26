package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_4(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		section  string
		expected map[string]string
		err      error
	}

	testCases := []testCase{
		{
			name:    "existing section",
			section: "section1",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			err: nil,
		},
		{
			name:    "non-existing section",
			section: "section3",
			expected: map[string]string{
				"key3": "value3",
			},
			err: errors.New("not exist section"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			container := &IniConfigContainer{
				data: map[string]map[string]string{
					"section1": {
						"key1": "value1",
						"key2": "value2",
					},
					"section2": {
						"key3": "value3",
					},
				},
			}

			result, err := container.GetSection(tc.section)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
			if err != nil && tc.err == nil {
				t.Errorf("Expected no error, got %v", err)
			} else if err == nil && tc.err != nil {
				t.Errorf("Expected error %v, got no error", tc.err)
			} else if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}
