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
			name: "Test case 1",
			item: Item{Type: tError},
			want: true,
		},
		{
			name: "Test case 2",
			item: Item{Type: tEOF},
			want: true,
		},
		{
			name: "Test case 3",
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
