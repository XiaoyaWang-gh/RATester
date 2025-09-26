package template

import (
	"fmt"
	"testing"
)

func TestevalArgs_2(t *testing.T) {
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
			args: []any{123, 456},
			want: "[123 456]",
		},
		{
			name: "Test case 3",
			args: []any{true, false},
			want: "[true false]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := evalArgs(tt.args); got != tt.want {
				t.Errorf("evalArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
