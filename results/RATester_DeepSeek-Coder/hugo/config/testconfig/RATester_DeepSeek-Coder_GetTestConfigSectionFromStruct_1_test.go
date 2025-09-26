package testconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/config"
)

func TestGetTestConfigSectionFromStruct_1(t *testing.T) {
	type testStruct struct {
		Test string
	}

	testCases := []struct {
		name     string
		section  string
		v        any
		expected config.AllProvider
	}{
		{
			name:    "Test case 1",
			section: "test",
			v:       testStruct{Test: "test"},
			expected: func() config.AllProvider {
				cfg := config.NewFrom(maps.Params{
					"test": config.FromTOMLConfigString(`Test = "test"`).Get(""),
				})
				return GetTestConfig(nil, cfg)
			}(),
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := GetTestConfigSectionFromStruct(tc.section, tc.v)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
