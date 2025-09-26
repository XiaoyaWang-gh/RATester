package pageparser

import (
	"fmt"
	"testing"
)

func TestIsShortcodeName_1(t *testing.T) {
	tests := []struct {
		name string
		item Item
		want bool
	}{
		{
			name: "Test when item type is tScName",
			item: Item{Type: tScName},
			want: true,
		},
		{
			name: "Test when item type is not tScName",
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

			if got := tt.item.IsShortcodeName(); got != tt.want {
				t.Errorf("IsShortcodeName() = %v, want %v", got, tt.want)
			}
		})
	}
}
