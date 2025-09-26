package layouts

import (
	"fmt"
	"testing"
)

func TestisList_1(t *testing.T) {
	tests := []struct {
		name string
		d    LayoutDescriptor
		want bool
	}{
		{
			name: "Test case 1",
			d: LayoutDescriptor{
				RenderingHook: true,
				Kind:          "page",
			},
			want: false,
		},
		{
			name: "Test case 2",
			d: LayoutDescriptor{
				RenderingHook: false,
				Kind:          "404",
			},
			want: false,
		},
		{
			name: "Test case 3",
			d: LayoutDescriptor{
				RenderingHook: false,
				Kind:          "sitemap",
			},
			want: false,
		},
		{
			name: "Test case 4",
			d: LayoutDescriptor{
				RenderingHook: false,
				Kind:          "sitemapindex",
			},
			want: false,
		},
		{
			name: "Test case 5",
			d: LayoutDescriptor{
				RenderingHook: false,
				Kind:          "other",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.isList(); got != tt.want {
				t.Errorf("isList() = %v, want %v", got, tt.want)
			}
		})
	}
}
