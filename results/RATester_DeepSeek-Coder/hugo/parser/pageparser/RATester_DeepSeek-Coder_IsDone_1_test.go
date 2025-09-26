package pageparser

import (
	"fmt"
	"testing"
)

func TestIsDone_1(t *testing.T) {
	tests := []struct {
		name string
		item Item
		want bool
	}{
		{
			name: "Test with Type as tError",
			item: Item{Type: tError},
			want: true,
		},
		{
			name: "Test with Type as tEOF",
			item: Item{Type: tEOF},
			want: true,
		},
		{
			name: "Test with Type as tText",
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

			if got := tt.item.IsDone(); got != tt.want {
				t.Errorf("IsDone() = %v, want %v", got, tt.want)
			}
		})
	}
}
