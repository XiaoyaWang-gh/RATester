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
			name: "Test case 1",
			m:    Type{SubType: "markdown"},
			want: true,
		},
		{
			name: "Test case 2",
			m:    Type{SubType: "html"},
			want: false,
		},
		{
			name: "Test case 3",
			m:    Type{SubType: "text"},
			want: false,
		},
		{
			name: "Test case 4",
			m:    Type{SubType: ""},
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
