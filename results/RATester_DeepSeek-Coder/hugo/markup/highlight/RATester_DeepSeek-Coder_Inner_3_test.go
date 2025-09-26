package highlight

import (
	"fmt"
	"html/template"
	"testing"
)

func TestInner_3(t *testing.T) {
	tests := []struct {
		name string
		h    HighlightResult
		want template.HTML
	}{
		{
			name: "Test case 1",
			h: HighlightResult{
				innerLow:    0,
				innerHigh:   5,
				highlighted: "<span>Hello</span>, world",
			},
			want: "<span>Hello</span>",
		},
		{
			name: "Test case 2",
			h: HighlightResult{
				innerLow:    0,
				innerHigh:   10,
				highlighted: "<span>Hello, world</span>",
			},
			want: "<span>Hello, world</span>",
		},
		{
			name: "Test case 3",
			h: HighlightResult{
				innerLow:    6,
				innerHigh:   11,
				highlighted: "<span>Hello</span>, world",
			},
			want: ", world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.h.Inner(); got != tt.want {
				t.Errorf("HighlightResult.Inner() = %v, want %v", got, tt.want)
			}
		})
	}
}
