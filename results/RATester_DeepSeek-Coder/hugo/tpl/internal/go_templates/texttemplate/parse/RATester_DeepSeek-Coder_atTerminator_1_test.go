package parse

import (
	"fmt"
	"testing"
)

func TestAtTerminator_1(t *testing.T) {
	tests := []struct {
		name string
		l    *lexer
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.l.atTerminator(); got != tt.want {
				t.Errorf("lexer.atTerminator() = %v, want %v", got, tt.want)
			}
		})
	}
}
