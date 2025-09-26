package middleware

import (
	"fmt"
	"net/url"
	"testing"
)

func TestRemoveTarget_1(t *testing.T) {
	b := &commonBalancer{
		targets: []*ProxyTarget{
			{Name: "target1", URL: &url.URL{}},
			{Name: "target2", URL: &url.URL{}},
			{Name: "target3", URL: &url.URL{}},
		},
	}

	tests := []struct {
		name     string
		target   string
		expected bool
	}{
		{"remove existing target", "target2", true},
		{"remove non-existing target", "target4", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := b.RemoveTarget(test.target)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
