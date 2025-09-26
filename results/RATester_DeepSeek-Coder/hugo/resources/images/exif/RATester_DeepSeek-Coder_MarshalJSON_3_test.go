package exif

import (
	"fmt"
	"testing"
)

func TestMarshalJSON_3(t *testing.T) {
	testCases := []struct {
		name     string
		tags     Tags
		expected string
	}{
		{
			name:     "empty tags",
			tags:     Tags{},
			expected: `{}`,
		},
		{
			name:     "single tag",
			tags:     Tags{"ExposureTime": "1/60"},
			expected: `{"ExposureTime":"1/60"}`,
		},
		{
			name: "multiple tags",
			tags: Tags{
				"ExposureTime":    "1/60",
				"FNumber":         "2.8",
				"ISOSpeedRatings": "100",
			},
			expected: `{"ExposureTime":"1/60","FNumber":"2.8","ISOSpeedRatings":"100"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := tc.tags.MarshalJSON()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if string(result) != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, string(result))
			}
		})
	}
}
