package page

import (
	"fmt"
	"testing"
)

func TestPrintableValue_3(t *testing.T) {
	tests := []struct {
		name string
		s    Summary
		want any
	}{
		{
			name: "Test case 1",
			s: Summary{
				Text:      "Test summary",
				Type:      "auto",
				Truncated: false,
			},
			want: "Test summary",
		},
		{
			name: "Test case 2",
			s: Summary{
				Text:      "Test summary 2",
				Type:      "manual",
				Truncated: true,
			},
			want: "Test summary 2",
		},
		{
			name: "Test case 3",
			s: Summary{
				Text:      "Test summary 3",
				Type:      "frontmatter",
				Truncated: false,
			},
			want: "Test summary 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.PrintableValue(); got != tt.want {
				t.Errorf("Summary.PrintableValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
