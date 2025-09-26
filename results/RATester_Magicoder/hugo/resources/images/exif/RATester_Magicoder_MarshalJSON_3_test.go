package exif

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_3(t *testing.T) {
	testCases := []struct {
		name     string
		input    Tags
		expected string
	}{
		{
			name: "empty tags",
			input: Tags{
				"": nil,
			},
			expected: `{}`,
		},
		{
			name: "single tag",
			input: Tags{
				"Exif.Image.Make": "Canon",
			},
			expected: `{"Exif.Image.Make":"Canon"}`,
		},
		{
			name: "multiple tags",
			input: Tags{
				"Exif.Image.Make":   "Canon",
				"Exif.Image.Model":  "Canon EOS 5D Mark IV",
				"Exif.Image.Orient": "1",
			},
			expected: `{"Exif.Image.Make":"Canon","Exif.Image.Model":"Canon EOS 5D Mark IV","Exif.Image.Orient":"1"}`,
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
				t.Fatalf("Unexpected error: %v", err)
			}

			if string(result) != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, string(result))
			}
		})
	}
}
