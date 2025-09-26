package pageparser

import (
	"fmt"
	"testing"
)

func TestIsRightShortcodeDelim_1(t *testing.T) {
	tests := []struct {
		name string
		i    Item
		want bool
	}{
		{
			name: "tRightDelimScWithMarkup",
			i: Item{
				Type: tRightDelimScWithMarkup,
			},
			want: true,
		},
		{
			name: "tRightDelimScNoMarkup",
			i: Item{
				Type: tRightDelimScNoMarkup,
			},
			want: true,
		},
		{
			name: "other",
			i: Item{
				Type: 10,
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

			if got := tt.i.IsRightShortcodeDelim(); got != tt.want {
				t.Errorf("Item.IsRightShortcodeDelim() = %v, want %v", got, tt.want)
			}
		})
	}
}
