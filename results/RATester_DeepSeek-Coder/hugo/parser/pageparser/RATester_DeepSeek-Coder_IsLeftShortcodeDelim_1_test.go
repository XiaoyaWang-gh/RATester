package pageparser

import (
	"fmt"
	"testing"
)

func TestIsLeftShortcodeDelim_1(t *testing.T) {
	tests := []struct {
		name string
		item Item
		want bool
	}{
		{
			name: "Test with left shortcode delimiter with markup",
			item: Item{Type: tLeftDelimScWithMarkup},
			want: true,
		},
		{
			name: "Test with left shortcode delimiter without markup",
			item: Item{Type: tLeftDelimScNoMarkup},
			want: true,
		},
		{
			name: "Test with other type",
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

			if got := tt.item.IsLeftShortcodeDelim(); got != tt.want {
				t.Errorf("IsLeftShortcodeDelim() = %v, want %v", got, tt.want)
			}
		})
	}
}
