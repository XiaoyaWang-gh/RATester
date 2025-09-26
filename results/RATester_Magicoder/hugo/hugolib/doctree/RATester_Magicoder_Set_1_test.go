package doctree

import (
	"fmt"
	"testing"
)

func TestSet_1(t *testing.T) {
	tests := []struct {
		name string
		d    DimensionFlag
		o    DimensionFlag
		want DimensionFlag
	}{
		{
			name: "Test case 1",
			d:    1,
			o:    2,
			want: 3,
		},
		{
			name: "Test case 2",
			d:    4,
			o:    8,
			want: 12,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.Set(tt.o); got != tt.want {
				t.Errorf("DimensionFlag.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}
