package template

import (
	"fmt"
	"testing"
)

func TestIsComment_1(t *testing.T) {
	type test struct {
		name string
		s    state
		want bool
	}

	tests := []test{
		{"stateHTMLCmt", stateHTMLCmt, true},
		{"stateJSBlockCmt", stateJSBlockCmt, true},
		{"stateJSLineCmt", stateJSLineCmt, true},
		{"stateJSHTMLOpenCmt", stateJSHTMLOpenCmt, true},
		{"stateCSSBlockCmt", stateCSSBlockCmt, true},
		{"stateNormal", state(0), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isComment(tt.s); got != tt.want {
				t.Errorf("isComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
