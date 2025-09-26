package hashing

import (
	"fmt"
	"testing"
)

func TestHashString_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{1, "2", 3.0, "4", true},
			want: "12345",
		},
		{
			name: "Test case 2",
			args: []any{},
			want: "0",
		},
		{
			name: "Test case 3",
			args: []any{nil},
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

			if got := HashString(tt.args...); got != tt.want {
				t.Errorf("HashString() = %v, want %v", got, tt.want)
			}
		})
	}
}
