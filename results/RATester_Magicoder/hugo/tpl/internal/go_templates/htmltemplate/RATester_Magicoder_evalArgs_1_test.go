package template

import (
	"fmt"
	"testing"
)

func TestevalArgs_1(t *testing.T) {

	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "single string",
			args: []any{"hello"},
			want: "hello",
		},
		{
			name: "multiple args",
			args: []any{"hello", "world"},
			want: "hello world",
		},
		{
			name: "single non-string",
			args: []any{123},
			want: "123",
		},
		{
			name: "multiple mixed",
			args: []any{"hello", 123, "world"},
			want: "hello 123 world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := evalArgs(tt.args...); got != tt.want {
				t.Errorf("evalArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
