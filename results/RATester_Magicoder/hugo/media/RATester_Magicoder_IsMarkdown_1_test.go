package media

import (
	"fmt"
	"testing"
)

func TestIsMarkdown_1(t *testing.T) {
	tests := []struct {
		name string
		m    Type
		want bool
	}{
		{
			name: "Markdown",
			m:    Type{SubType: Builtin.MarkdownType.SubType},
			want: true,
		},
		{
			name: "Not Markdown",
			m:    Type{SubType: "notMarkdown"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.IsMarkdown(); got != tt.want {
				t.Errorf("IsMarkdown() = %v, want %v", got, tt.want)
			}
		})
	}
}
