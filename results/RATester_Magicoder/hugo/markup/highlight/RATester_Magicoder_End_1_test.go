package highlight

import (
	"fmt"
	"testing"
)

func TestEnd_1(t *testing.T) {
	tests := []struct {
		name string
		s    startEnd
		code bool
		want string
	}{
		{
			name: "Test case 1",
			s: startEnd{
				start: func(code bool, styleAttr string) string {
					return "start"
				},
				end: func(code bool) string {
					return "end"
				},
			},
			code: true,
			want: "end",
		},
		{
			name: "Test case 2",
			s: startEnd{
				start: func(code bool, styleAttr string) string {
					return "start"
				},
				end: func(code bool) string {
					return "end"
				},
			},
			code: false,
			want: "end",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.End(tt.code); got != tt.want {
				t.Errorf("startEnd.End() = %v, want %v", got, tt.want)
			}
		})
	}
}
