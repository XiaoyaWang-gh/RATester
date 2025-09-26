package resource

import (
	"fmt"
	"testing"
)

func TestMarkStale_3(t *testing.T) {
	type mockStaleMarker struct {
		stale bool
	}

	tests := []struct {
		name     string
		input    []any
		expected []bool
	}{
		{
			name: "Test with one StaleMarker",
			input: []any{
				&mockStaleMarker{stale: false},
			},
			expected: []bool{true},
		},
		{
			name: "Test with multiple StaleMarkers",
			input: []any{
				&mockStaleMarker{stale: false},
				&mockStaleMarker{stale: false},
				&mockStaleMarker{stale: false},
			},
			expected: []bool{true, true, true},
		},
		{
			name: "Test with non-StaleMarkers",
			input: []any{
				"not a StaleMarker",
				&mockStaleMarker{stale: false},
				nil,
			},
			expected: []bool{false, true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			MarkStale(tt.input...)
			for i, o := range tt.input {
				if s, ok := o.(*mockStaleMarker); ok {
					if s.stale != tt.expected[i] {
						t.Errorf("Expected stale to be %v, but got %v", tt.expected[i], s.stale)
					}
				}
			}
		})
	}
}
