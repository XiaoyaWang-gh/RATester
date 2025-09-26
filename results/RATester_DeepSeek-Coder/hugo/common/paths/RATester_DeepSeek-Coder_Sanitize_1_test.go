package paths

import (
	"fmt"
	"testing"
)

func TestSanitize_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Test case 1",
			arg:  "Hello, World",
			want: "Hello-World",
		},
		{
			name: "Test case 2",
			arg:  "Hello,  World",
			want: "Hello-World",
		},
		{
			name: "Test case 3",
			arg:  "Hello,   World",
			want: "Hello-World",
		},
		{
			name: "Test case 4",
			arg:  "Hello,    World",
			want: "Hello-World",
		},
		{
			name: "Test case 5",
			arg:  "Hello,     World",
			want: "Hello-World",
		},
		{
			name: "Test case 6",
			arg:  "Hello,      World",
			want: "Hello-World",
		},
		{
			name: "Test case 7",
			arg:  "Hello,       World",
			want: "Hello-World",
		},
		{
			name: "Test case 8",
			arg:  "Hello,        World",
			want: "Hello-World",
		},
		{
			name: "Test case 9",
			arg:  "Hello,         World",
			want: "Hello-World",
		},
		{
			name: "Test case 10",
			arg:  "Hello,          World",
			want: "Hello-World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Sanitize(tt.arg); got != tt.want {
				t.Errorf("Sanitize() = %v, want %v", got, tt.want)
			}
		})
	}
}
