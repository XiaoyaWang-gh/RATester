package file

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetDefaults_19(t *testing.T) {
	type testCase struct {
		name     string
		provider *Provider
		expected *Provider
	}

	testCases := []testCase{
		{
			name: "Test SetDefaults",
			provider: &Provider{
				Directory: "test",
				Watch:     false,
				Filename:  "test.yaml",
			},
			expected: &Provider{
				Directory: "test",
				Watch:     true,
				Filename:  "",
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

			tc.provider.SetDefaults()

			if !reflect.DeepEqual(tc.provider, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, tc.provider)
			}
		})
	}
}
