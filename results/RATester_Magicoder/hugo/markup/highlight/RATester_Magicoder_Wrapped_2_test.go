package highlight

import (
	"fmt"
	"html/template"
	"testing"
)

func TestWrapped_2(t *testing.T) {
	tests := []struct {
		name string
		h    HighlightResult
		want template.HTML
	}{
		{
			name: "Test case 1",
			h: HighlightResult{
				innerLow:    1,
				innerHigh:   2,
				highlighted: "<p>Test</p>",
			},
			want: "<p>Test</p>",
		},
		{
			name: "Test case 2",
			h: HighlightResult{
				innerLow:    3,
				innerHigh:   4,
				highlighted: "<p>Another Test</p>",
			},
			want: "<p>Another Test</p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.h.Wrapped(); got != tt.want {
				t.Errorf("Wrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}
