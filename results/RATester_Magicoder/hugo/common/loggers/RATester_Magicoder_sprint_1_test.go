package loggers

import (
	"fmt"
	"testing"
)

func Testsprint_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test with single string",
			args: []any{"Hello, World!"},
			want: "Hello, World!",
		},
		{
			name: "Test with multiple strings",
			args: []any{"Hello,", "World!"},
			want: "Hello, World!",
		},
		{
			name: "Test with mixed types",
			args: []any{"Hello,", 123, "World!"},
			want: "Hello, 123 World!",
		},
		{
			name: "Test with no arguments",
			args: []any{},
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

			l := &logAdapter{}
			if got := l.sprint(tt.args...); got != tt.want {
				t.Errorf("sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
