package modules

import (
	"fmt"
	"testing"
)

func TestComponentAndName_1(t *testing.T) {
	type testCase struct {
		name     string
		mount    Mount
		expected struct {
			component string
			name      string
		}
	}

	testCases := []testCase{
		{
			name: "simple",
			mount: Mount{
				Target: "assets/bootstrap/scss",
			},
			expected: struct {
				component string
				name      string
			}{
				component: "assets",
				name:      "bootstrap/scss",
			},
		},
		{
			name: "complex",
			mount: Mount{
				Target: "src/components/bootstrap/scss",
			},
			expected: struct {
				component string
				name      string
			}{
				component: "src/components",
				name:      "bootstrap/scss",
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

			component, name := tc.mount.ComponentAndName()
			if component != tc.expected.component || name != tc.expected.name {
				t.Errorf("Expected (%s, %s), got (%s, %s)", tc.expected.component, tc.expected.name, component, name)
			}
		})
	}
}
