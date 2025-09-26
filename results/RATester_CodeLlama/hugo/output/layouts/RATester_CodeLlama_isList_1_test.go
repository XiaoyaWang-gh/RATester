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
			name: "not list",
			d: LayoutDescriptor{
				RenderingHook: false,
				Kind:          "page",
			},
			want: false,
		},
		{
			name: "list",
			d: LayoutDescriptor{
				RenderingHook: false,
				Kind:          "list",
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
