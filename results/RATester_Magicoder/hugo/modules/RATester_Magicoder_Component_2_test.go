package modules

import (
	"fmt"
	"testing"
)

func TestComponent_2(t *testing.T) {
	tests := []struct {
		name     string
		mount    Mount
		expected string
	}{
		{
			name: "TestComponent_0",
			mount: Mount{
				Source: "source",
				Target: "target",
			},
			expected: "target",
		},
		{
			name: "TestComponent_1",
			mount: Mount{
				Source: "source",
				Target: "target/sub",
			},
			expected: "target",
		},
		{
			name: "TestComponent_2",
			mount: Mount{
				Source: "source",
				Target: "target/sub/sub",
			},
			expected: "target",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.mount.Component()
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
