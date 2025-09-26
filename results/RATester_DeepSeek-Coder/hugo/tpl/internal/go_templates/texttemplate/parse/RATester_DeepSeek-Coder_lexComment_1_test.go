package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLexComment_1(t *testing.T) {
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

			if got := lexComment(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexComment() = %v, want %v", got, tt.want)
			}
		})
	}
}
