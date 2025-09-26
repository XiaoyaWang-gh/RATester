package loggers

import (
	"fmt"
	"strings"
	"testing"
)

func TestErrors_2(t *testing.T) {
	tests := []struct {
		name string
		l    *logAdapter
		want string
	}{
		{
			name: "Test case 1",
			l: &logAdapter{
				errors: &strings.Builder{},
			},
			want: "",
		},
		{
			name: "Test case 2",
			l: &logAdapter{
				errors: func() *strings.Builder {
					sb := &strings.Builder{}
					sb.WriteString("Error 1")
					return sb
				}(),
			},
			want: "Error 1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.l.Errors(); got != tt.want {
				t.Errorf("Errors() = %v, want %v", got, tt.want)
			}
		})
	}
}
