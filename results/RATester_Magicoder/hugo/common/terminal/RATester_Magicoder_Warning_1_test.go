package terminal

import (
	"fmt"
	"testing"
)

func TestWarning_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test case 1",
			args: "test",
			want: "warning: test",
		},
		{
			name: "Test case 2",
			args: "test2",
			want: "warning: test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Warning(tt.args); got != tt.want {
				t.Errorf("Warning() = %v, want %v", got, tt.want)
			}
		})
	}
}
