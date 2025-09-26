package utils

import (
	"fmt"
	"testing"
)

func TestInSlice_1(t *testing.T) {
	tests := []struct {
		name string
		v    string
		sl   []string
		want bool
	}{
		{
			name: "Test case 1",
			v:    "test",
			sl:   []string{"test", "test1", "test2"},
			want: true,
		},
		{
			name: "Test case 2",
			v:    "test3",
			sl:   []string{"test", "test1", "test2"},
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

			if got := InSlice(tt.v, tt.sl); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
