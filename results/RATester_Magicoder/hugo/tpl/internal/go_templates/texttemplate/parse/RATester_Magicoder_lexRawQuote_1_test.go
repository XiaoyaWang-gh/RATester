package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlexRawQuote_1(t *testing.T) {
	tests := []struct {
		name string
		l    *lexer
		want stateFn
	}{
		{
			name: "test case 1",
			l: &lexer{
				input: "`test`",
			},
			want: func(l *lexer) stateFn {
				return l.emit(itemRawString)
			},
		},
		{
			name: "test case 2",
			l: &lexer{
				input: "`test",
			},
			want: func(l *lexer) stateFn {
				return l.errorf("unterminated raw quoted string")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := lexRawQuote(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexRawQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
