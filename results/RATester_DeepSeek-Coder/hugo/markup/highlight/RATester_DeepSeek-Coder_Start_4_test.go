package highlight

import (
	"fmt"
	"testing"
)

func TestStart_4(t *testing.T) {
	tests := []struct {
		name      string
		s         startEnd
		code      bool
		styleAttr string
		want      string
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
			code:      true,
			styleAttr: "style",
			want:      "start",
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
			code:      false,
			styleAttr: "style",
			want:      "start",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.Start(tt.code, tt.styleAttr); got != tt.want {
				t.Errorf("startEnd.Start() = %v, want %v", got, tt.want)
			}
		})
	}
}
