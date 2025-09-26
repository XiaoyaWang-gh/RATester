package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLexInsideAction_1(t *testing.T) {
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

			if got := lexInsideAction(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexInsideAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
