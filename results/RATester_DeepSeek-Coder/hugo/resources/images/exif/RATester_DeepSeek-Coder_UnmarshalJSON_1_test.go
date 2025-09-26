package exif

import (
	"errors"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestUnmarshalJSON_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    string
		expected Tags
		err      error
	}{
		{
			name:  "valid json",
			input: `{"Make": "Canon", "Model": "Canon EOS 5D Mark IV", "DateTime": "2022:01:02 15:04:05"}`,
			expected: Tags{
				"Make":     "Canon",
				"Model":    "Canon EOS 5D Mark IV",
				"DateTime": "2022:01:02 15:04:05",
			},
			err: nil,
		},
		{
			name:     "invalid json",
			input:    `{"Make": "Canon", "Model": "Canon EOS 5D Mark IV", "DateTime": "2022:01:02 15:04:05",}`,
			expected: nil,
			err:      errors.New("unexpected end of JSON input"),
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			var v Tags
			err := v.UnmarshalJSON([]byte(tc.input))

			if tc.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, v)
			}
		})
	}
}
