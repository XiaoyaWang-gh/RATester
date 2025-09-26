package template

import (
	"fmt"
	"testing"
)

func TestIsInScriptLiteral_1(t *testing.T) {
	tests := []struct {
		name string
		s    state
		want bool
	}{
		{"stateJSDqStr", stateJSDqStr, true},
		{"stateJSSqStr", stateJSSqStr, true},
		{"stateJSTmplLit", stateJSTmplLit, true},
		{"stateJSRegexp", stateJSRegexp, true},
		{"non-script-literal state", 0, false},
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
