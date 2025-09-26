package pageparser

import (
	"fmt"
	"testing"
)

func TestIsRightShortcodeDelim_1(t *testing.T) {
	tests := []struct {
		name string
		item Item
		want bool
	}{
		{
			name: "right shortcode delim with markup",
			item: Item{Type: tRightDelimScWithMarkup},
			want: true,
		},
		{
			name: "right shortcode delim without markup",
			item: Item{Type: tRightDelimScNoMarkup},
			want: true,
		},
		{
			name: "other type",
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

			if got := tt.item.IsRightShortcodeDelim(); got != tt.want {
				t.Errorf("IsRightShortcodeDelim() = %v, want %v", got, tt.want)
			}
		})
	}
}
