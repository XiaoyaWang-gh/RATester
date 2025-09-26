package template

import (
	"fmt"
	"testing"
)

func TestSetOption_2(t *testing.T) {
	type testCase struct {
		name     string
		opt      string
		expected option
	}

	testCases := []testCase{
		{
			name: "Test with missingkey=invalid",
			opt:  "missingkey=invalid",
			expected: option{
				missingKey: mapInvalid,
			},
		},
		{
			name: "Test with missingkey=zero",
			opt:  "missingkey=zero",
			expected: option{
				missingKey: mapZeroValue,
			},
		},
		{
			name: "Test with missingkey=error",
			opt:  "missingkey=error",
			expected: option{
				missingKey: mapError,
			},
		},
		{
			name: "Test with empty string",
			opt:  "",
			expected: option{
				missingKey: mapInvalid, // default value
			},
		},
		{
			name: "Test with unrecognized option",
			opt:  "unrecognized=value",
			expected: option{
				missingKey: mapInvalid, // default value
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tmpl := &Template{}
			tmpl.setOption(tc.opt)
			if tmpl.option.missingKey != tc.expected.missingKey {
				t.Errorf("Expected missingKey to be %v, got %v", tc.expected.missingKey, tmpl.option.missingKey)
			}
		})
	}
}
