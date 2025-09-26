package tcp

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseHostSNI_1(t *testing.T) {
	type testCase struct {
		name     string
		rule     string
		expected []string
		err      error
	}

	testCases := []testCase{
		{
			name:     "valid rule",
			rule:     "HostSNI(`example.com`)",
			expected: []string{"example.com"},
			err:      nil,
		},
		{
			name:     "invalid rule",
			rule:     "HostSNI(`example.com`",
			expected: nil,
			err:      fmt.Errorf("error while parsing rule HostSNI(`example.com`"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ParseHostSNI(tc.rule)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
			if err != nil && tc.err == nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if err == nil && tc.err != nil {
				t.Errorf("Expected error %v, got no error", tc.err)
			}
			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}
		})
	}
}
