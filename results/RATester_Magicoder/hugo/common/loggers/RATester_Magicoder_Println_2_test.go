package loggers

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintln_2(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test with string",
			args: []any{"Hello", "World"},
			want: "Hello World\n",
		},
		{
			name: "Test with int",
			args: []any{1, 2, 3},
			want: "1 2 3\n",
		},
		{
			name: "Test with float",
			args: []any{1.1, 2.2, 3.3},
			want: "1.1 2.2 3.3\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			out := &strings.Builder{}
			la := &logAdapter{out: out}
			la.Println(tt.args...)
			got := out.String()
			if got != tt.want {
				t.Errorf("Println() = %v, want %v", got, tt.want)
			}
		})
	}
}
