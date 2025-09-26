package config

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDefaultFloat_3(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		key        string
		defaultVal float64
		want       float64
	}

	testCases := []testCase{
		{
			name:       "TestDefaultFloat_ExistingKey",
			key:        "existingKey",
			defaultVal: 1.0,
			want:       2.0,
		},
		{
			name:       "TestDefaultFloat_NonExistingKey",
			key:        "nonExistingKey",
			defaultVal: 3.0,
			want:       3.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &IniConfigContainer{
				data: map[string]map[string]string{
					"section": {
						tc.key: strconv.FormatFloat(tc.want, 'f', -1, 64),
					},
				},
			}

			got := c.DefaultFloat(tc.key, tc.defaultVal)
			if got != tc.want {
				t.Errorf("DefaultFloat() = %v, want %v", got, tc.want)
			}
		})
	}
}
