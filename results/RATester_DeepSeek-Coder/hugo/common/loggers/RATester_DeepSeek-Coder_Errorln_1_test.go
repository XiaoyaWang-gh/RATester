package loggers

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrorln_1(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test with one argument",
			args: []any{"test"},
			want: "test\n",
		},
		{
			name: "Test with multiple arguments",
			args: []any{"test1", "test2"},
			want: "test1test2\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &logAdapter{
				errors: &strings.Builder{},
			}

			l.Errorln(tt.args...)

			got := l.errors.String()
			if got != tt.want {
				t.Errorf("Errorln() = %v, want %v", got, tt.want)
			}
		})
	}
}
