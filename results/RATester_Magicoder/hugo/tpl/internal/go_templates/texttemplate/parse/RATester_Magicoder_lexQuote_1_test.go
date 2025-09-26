package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlexQuote_1(t *testing.T) {
	tests := []struct {
		name string
		l    *lexer
		want stateFn
	}{
		{
			name: "test case 1",
			l: &lexer{
				input: "\"test string\"",
			},
			want: func(l *lexer) stateFn {
				return l.emit(itemString)
			},
		},
		{
			name: "test case 2",
			l: &lexer{
				input: "\"test string",
			},
			want: func(l *lexer) stateFn {
				return l.errorf("unterminated quoted string")
			},
		},
		{
			name: "test case 3",
			l: &lexer{
				input: "\"test string\n",
			},
			want: func(l *lexer) stateFn {
				return l.errorf("unterminated quoted string")
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

			if got := lexQuote(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}
