package pageparser

import (
	"fmt"
	"testing"
)

func TestIsShortcodeMarkupDelimiter_1(t *testing.T) {
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
			name: "tRightDelimScWithMarkup",
			i: Item{
				Type: tRightDelimScWithMarkup,
			},
			want: true,
		},
		{
			name: "other",
			i: Item{
				Type: 12,
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

			if got := tt.i.IsShortcodeMarkupDelimiter(); got != tt.want {
				t.Errorf("Item.IsShortcodeMarkupDelimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}
