package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlexNumber_1(t *testing.T) {
	tests := []struct {
		name string
		l    *lexer
		want stateFn
	}{
		{
			name: "Test case 1",
			l: &lexer{
				input: "123",
			},
			want: func(l *lexer) stateFn {
				if !l.scanNumber() {
					return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
				}
				return l.emit(itemNumber)
			},
		},
		{
			name: "Test case 2",
			l: &lexer{
				input: "123+456i",
			},
			want: func(l *lexer) stateFn {
				if !l.scanNumber() {
					return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
				}
				if sign := l.peek(); sign == '+' || sign == '-' {

					if !l.scanNumber() || l.input[l.pos-1] != 'i' {
						return l.errorf("bad number syntax: %q", l.input[l.start:l.pos])
					}
					return l.emit(itemComplex)
				}
				return l.emit(itemNumber)
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

			if got := lexNumber(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
