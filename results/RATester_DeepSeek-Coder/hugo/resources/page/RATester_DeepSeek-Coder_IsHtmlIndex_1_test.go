package page

import (
	"fmt"
	"testing"
)

func TestIsHtmlIndex_1(t *testing.T) {
	tests := []struct {
		name string
		p    *pagePathBuilder
		want bool
	}{
		{
			name: "Test when last element is 'index.html'",
			p: &pagePathBuilder{
				els: []string{"test.html", "index.html"},
			},
			want: true,
		},
		{
			name: "Test when last element is not 'index.html'",
			p: &pagePathBuilder{
				els: []string{"test.html", "test2.html"},
			},
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

			if got := tt.p.IsHtmlIndex(); got != tt.want {
				t.Errorf("pagePathBuilder.IsHtmlIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
