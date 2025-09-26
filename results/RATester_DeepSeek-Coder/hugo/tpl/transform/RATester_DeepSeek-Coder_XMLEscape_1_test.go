package transform

import (
	"errors"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestXMLEscape_1(t *testing.T) {
	type testCase struct {
		name     string
		input    any
		expected string
		err      error
	}

	testCases := []testCase{
		{
			name:     "valid string",
			input:    "<tag>content</tag>",
			expected: "&lt;tag&gt;content&lt;/tag&gt;",
			err:      nil,
		},
		{
			name:     "valid string with special characters",
			input:    "<tag>content & special characters</tag>",
			expected: "&lt;tag&gt;content &amp; special characters&lt;/tag&gt;",
			err:      nil,
		},
		{
			name:     "invalid type",
			input:    123,
			expected: "",
			err:      errors.New("cast: unable to cast int to string"),
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
			err:      nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ns := &Namespace{}
			result, err := ns.XMLEscape(tc.input)
			if tc.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expected, result)
		})
	}
}
