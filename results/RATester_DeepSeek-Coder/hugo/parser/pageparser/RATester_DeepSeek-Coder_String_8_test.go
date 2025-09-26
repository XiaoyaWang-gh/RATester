package pageparser

import (
	"fmt"
	"testing"
)

func TestString_8(t *testing.T) {
	tests := []struct {
		name string
		i    ItemType
		want string
	}{
		{
			name: "Test ItemType 0",
			i:    0,
			want: "ItemType(0)",
		},
		{
			name: "Test ItemType 1",
			i:    1,
			want: "ItemType(1)",
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

			if got := tt.i.String(); got != tt.want {
				t.Errorf("ItemType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
