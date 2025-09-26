package logs

import (
	"fmt"
	"testing"
)

func TestMsgFunc_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case 1",
			args: []any{"Hello", "World"},
			want: "Hello World",
		},
		{
			name: "Test case 2",
			args: []any{1, 2, 3},
			want: "1 2 3",
		},
		{
			name: "Test case 3",
			args: []any{true, false},
			want: "true false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := msgFunc(tt.args...); got() != tt.want {
				t.Errorf("msgFunc() = %v, want %v", got(), tt.want)
			}
		})
	}
}
