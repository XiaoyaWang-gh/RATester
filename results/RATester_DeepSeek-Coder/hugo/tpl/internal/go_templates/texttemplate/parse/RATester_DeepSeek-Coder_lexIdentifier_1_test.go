package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLexIdentifier_1(t *testing.T) {
	tests := []struct {
		name string
		l    *lexer
		want stateFn
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

			if got := lexIdentifier(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}
