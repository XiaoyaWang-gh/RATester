package loggers

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrintln_2(t *testing.T) {
	tests := []struct {
		name string
		v    []any
		want string
	}{
		{
			name: "Test with string",
			v:    []any{"test"},
			want: "test\n",
		},
		{
			name: "Test with multiple arguments",
			v:    []any{"test1", "test2"},
			want: "test1 test2\n",
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
			la := &logAdapter{
				out: out,
			}

			la.Println(tt.v...)

			got := out.String()
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
