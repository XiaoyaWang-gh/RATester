package highlight

import (
	"fmt"
	"testing"
)

func TestStart_4(t *testing.T) {
	start := func(code bool, styleAttr string) string {
		return "start"
	}
	end := func(code bool) string {
		return "end"
	}

	s := startEnd{
		start: start,
		end:   end,
	}

	tests := []struct {
		name      string
		code      bool
		styleAttr string
		want      string
	}{
		{
			name:      "Test case 1",
			code:      true,
			styleAttr: "style",
			want:      "start",
		},
		{
			name:      "Test case 2",
			code:      false,
			styleAttr: "style",
			want:      "end",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := s.Start(tt.code, tt.styleAttr); got != tt.want {
				t.Errorf("Start() = %v, want %v", got, tt.want)
			}
		})
	}
}
