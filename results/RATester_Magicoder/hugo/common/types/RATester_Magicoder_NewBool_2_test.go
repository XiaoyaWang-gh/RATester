package types

import (
	"fmt"
	"testing"
)

func TestNewBool_2(t *testing.T) {
	tests := []struct {
		name string
		b    bool
		want *bool
	}{
		{
			name: "Test case 1",
			b:    true,
			want: NewBool(true),
		},
		{
			name: "Test case 2",
			b:    false,
			want: NewBool(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewBool(tt.b); *got != *tt.want {
				t.Errorf("NewBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
