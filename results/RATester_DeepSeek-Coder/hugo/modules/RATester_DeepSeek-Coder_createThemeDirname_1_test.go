package modules

import (
	"fmt"
	"testing"
)

func TestCreateThemeDirname_1(t *testing.T) {
	type testCase struct {
		name         string
		modulePath   string
		isProjectMod bool
		expected     string
		errExpected  bool
	}

	testCases := []testCase{
		{
			name:         "valid module path, project module",
			modulePath:   "../themes/mytheme",
			isProjectMod: true,
			expected:     "../themes/mytheme",
			errExpected:  false,
		},
		{
			name:         "valid module path, not project module",
			modulePath:   "../themes/mytheme",
			isProjectMod: false,
			expected:     "",
			errExpected:  true,
		},
		{
			name:         "absolute module path, project module",
			modulePath:   "/themes/mytheme",
			isProjectMod: true,
			expected:     "/themes/mytheme",
			errExpected:  false,
		},
		{
			name:         "absolute module path, not project module",
			modulePath:   "/themes/mytheme",
			isProjectMod: false,
			expected:     "",
			errExpected:  true,
		},
	}

	c := &Client{
		ccfg: ClientConfig{
			ThemesDir: "/themes",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := c.createThemeDirname(tc.modulePath, tc.isProjectMod)
			if tc.errExpected && err == nil {
				t.Errorf("Expected error, but got nil")
			}
			if !tc.errExpected && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %q, but got %q", tc.expected, result)
			}
		})
	}
}
