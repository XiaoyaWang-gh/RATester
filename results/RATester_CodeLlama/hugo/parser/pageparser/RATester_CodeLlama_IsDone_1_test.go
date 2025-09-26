package pageparser

import (
	"fmt"
	"testing"
)

func TestIsDone_1(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		i    Item
		want bool
	}{
		{
			name: "tError",
			i: Item{
				Type: tError,
			},
			want: true,
		},
		{
			name: "tEOF",
			i: Item{
				Type: tEOF,
			},
			want: true,
		},
		{
			name: "other",
			i: Item{
				Type: tError,
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

			if got := tt.i.IsDone(); got != tt.want {
				t.Errorf("Item.IsDone() = %v, want %v", got, tt.want)
			}
		})
	}
}
