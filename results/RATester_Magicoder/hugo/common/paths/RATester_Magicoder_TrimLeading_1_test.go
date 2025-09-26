package paths

import (
	"fmt"
	"testing"
)

func TestTrimLeading_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test case 1",
			args: "/test",
			want: "test",
		},
		{
			name: "Test case 2",
			args: "test",
			want: "test",
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

			if got := TrimLeading(tt.args); got != tt.want {
				t.Errorf("TrimLeading() = %v, want %v", got, tt.want)
			}
		})
	}
}
