package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlexFieldOrVariable_1(t *testing.T) {
	tests := []struct {
		name string
		l    *lexer
		typ  itemType
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

			if got := lexFieldOrVariable(tt.l, tt.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexFieldOrVariable() = %v, want %v", got, tt.want)
			}
		})
	}
}
