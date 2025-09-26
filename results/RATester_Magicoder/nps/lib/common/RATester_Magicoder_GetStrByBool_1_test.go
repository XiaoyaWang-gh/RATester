package common

import (
	"fmt"
	"testing"
)

func TestGetStrByBool_1(t *testing.T) {
	tests := []struct {
		name string
		b    bool
		want string
	}{
		{
			name: "Test case 1",
			b:    true,
			want: "1",
		},
		{
			name: "Test case 2",
			b:    false,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetStrByBool(tt.b); got != tt.want {
				t.Errorf("GetStrByBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
