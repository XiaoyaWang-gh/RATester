package pageparser

import (
	"fmt"
	"testing"
)

func TestIsLeftShortcodeDelim_1(t *testing.T) {
	tests := []struct {
		name string
		i    Item
		want bool
	}{
		{
			name: "tLeftDelimScWithMarkup",
			i: Item{
				Type: tLeftDelimScWithMarkup,
			},
			want: true,
		},
		{
			name: "tLeftDelimScNoMarkup",
			i: Item{
				Type: tLeftDelimScNoMarkup,
			},
			want: true,
		},
		{
			name: "other",
			i: Item{
				Type: 1,
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

			if got := tt.i.IsLeftShortcodeDelim(); got != tt.want {
				t.Errorf("Item.IsLeftShortcodeDelim() = %v, want %v", got, tt.want)
			}
		})
	}
}
