package pageparser

import (
	"fmt"
	"testing"
)

func TestIsFrontMatter_1(t *testing.T) {
	tests := []struct {
		name string
		item Item
		want bool
	}{
		{
			name: "Test with TypeFrontMatterYAML",
			item: Item{Type: TypeFrontMatterYAML},
			want: true,
		},
		{
			name: "Test with TypeFrontMatterORG",
			item: Item{Type: TypeFrontMatterORG},
			want: true,
		},
		{
			name: "Test with TypeFrontMatterTOML",
			item: Item{Type: TypeFrontMatterTOML},
			want: false,
		},
		{
			name: "Test with TypeFrontMatterJSON",
			item: Item{Type: TypeFrontMatterJSON},
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

			if got := tt.item.IsFrontMatter(); got != tt.want {
				t.Errorf("IsFrontMatter() = %v, want %v", got, tt.want)
			}
		})
	}
}
