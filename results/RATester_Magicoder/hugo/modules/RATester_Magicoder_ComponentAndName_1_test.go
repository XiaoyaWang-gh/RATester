package modules

import (
	"fmt"
	"testing"
)

func TestComponentAndName_1(t *testing.T) {
	type test struct {
		name     string
		mount    Mount
		expected struct {
			component string
			name      string
		}
	}

	tests := []test{
		{
			name: "simple",
			mount: Mount{
				Target: "assets/bootstrap/scss",
			},
			expected: struct {
				component string
				name      string
			}{
				component: "assets/bootstrap",
				name:      "scss",
			},
		},
		{
			name: "complex",
			mount: Mount{
				Target: "assets/bootstrap/scss/components/buttons",
			},
			expected: struct {
				component string
				name      string
			}{
				component: "assets/bootstrap/scss/components",
				name:      "buttons",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			component, name := tt.mount.ComponentAndName()
			if component != tt.expected.component {
				t.Errorf("expected component %q, got %q", tt.expected.component, component)
			}
			if name != tt.expected.name {
				t.Errorf("expected name %q, got %q", tt.expected.name, name)
			}
		})
	}
}
