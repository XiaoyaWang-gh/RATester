package paths

import (
	"fmt"
	"testing"
)

func TestTrimTrailing_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test case 1",
			args: "hello/",
			want: "hello",
		},
		{
			name: "Test case 2",
			args: "world",
			want: "world",
		},
		{
			name: "Test case 3",
			args: "/",
			want: "",
		},
		{
			name: "Test case 4",
			args: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := TrimTrailing(tt.args); got != tt.want {
				t.Errorf("TrimTrailing() = %v, want %v", got, tt.want)
			}
		})
	}
}
