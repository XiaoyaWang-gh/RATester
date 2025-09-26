package template

import (
	"fmt"
	"testing"
)

func TestisInScriptLiteral_1(t *testing.T) {

	tests := []struct {
		name string
		s    state
		want bool
	}{
		{
			name: "Test case 1",
			s:    stateJSDqStr,
			want: true,
		},
		{
			name: "Test case 2",
			s:    stateJSSqStr,
			want: true,
		},
		{
			name: "Test case 3",
			s:    stateJSTmplLit,
			want: true,
		},
		{
			name: "Test case 4",
			s:    0,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isInScriptLiteral(tt.s); got != tt.want {
				t.Errorf("isInScriptLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}
