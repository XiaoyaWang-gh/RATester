package terminal

import (
	"fmt"
	"testing"
)

func TestNotice_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test Notice",
			args: "Hello, World!",
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Notice(tt.args); got != tt.want {
				t.Errorf("Notice() = %v, want %v", got, tt.want)
			}
		})
	}
}
