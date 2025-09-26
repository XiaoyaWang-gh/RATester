package layouts

import (
	"fmt"
	"testing"
)

func TestIsList_1(t *testing.T) {
	tests := []struct {
		name string
		d    LayoutDescriptor
		want bool
	}{
		{
			name: "Test 1",
			d: LayoutDescriptor{
				Kind: "page",
			},
			want: false,
		},
		{
			name: "Test 2",
			d: LayoutDescriptor{
				Kind: "404",
			},
			want: false,
		},
		{
			name: "Test 3",
			d: LayoutDescriptor{
				Kind: "sitemap",
			},
			want: false,
		},
		{
			name: "Test 4",
			d: LayoutDescriptor{
				Kind: "sitemapindex",
			},
			want: false,
		},
		{
			name: "Test 5",
			d: LayoutDescriptor{
				Kind: "other",
			},
			want: true,
		},
		{
			name: "Test 6",
			d: LayoutDescriptor{
				RenderingHook: true,
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
				t.Errorf("LayoutDescriptor.isList() = %v, want %v", got, tt.want)
			}
		})
	}
}
