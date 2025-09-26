package pageparser

import (
	"fmt"
	"testing"
)

func TestIsIndentation_1(t *testing.T) {
	tests := []struct {
		name string
		item Item
		want bool
	}{
		{
			name: "Test with indentation item",
			item: Item{Type: tIndentation},
			want: true,
		},
		{
			name: "Test with non-indentation item",
			item: Item{Type: tText},
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

			if got := tt.item.IsIndentation(); got != tt.want {
				t.Errorf("IsIndentation() = %v, want %v", got, tt.want)
			}
		})
	}
}
