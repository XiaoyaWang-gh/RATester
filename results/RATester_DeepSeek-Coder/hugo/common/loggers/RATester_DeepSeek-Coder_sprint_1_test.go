package loggers

import (
	"fmt"
	"testing"
)

func TestSprint_1(t *testing.T) {
	tests := []struct {
		name string
		v    []any
		want string
	}{
		{
			name: "Test with string",
			v:    []any{"test"},
			want: "test",
		},
		{
			name: "Test with multiple strings",
			v:    []any{"test1", "test2"},
			want: "test1 test2",
		},
		{
			name: "Test with int",
			v:    []any{1},
			want: "1",
		},
		{
			name: "Test with multiple ints",
			v:    []any{1, 2},
			want: "1 2",
		},
		{
			name: "Test with float",
			v:    []any{1.1},
			want: "1.1",
		},
		{
			name: "Test with multiple floats",
			v:    []any{1.1, 2.2},
			want: "1.1 2.2",
		},
		{
			name: "Test with bool",
			v:    []any{true},
			want: "true",
		},
		{
			name: "Test with multiple bools",
			v:    []any{true, false},
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

			l := &logAdapter{}
			if got := l.sprint(tt.v...); got != tt.want {
				t.Errorf("sprint() = %v, want %v", got, tt.want)
			}
		})
	}
}
