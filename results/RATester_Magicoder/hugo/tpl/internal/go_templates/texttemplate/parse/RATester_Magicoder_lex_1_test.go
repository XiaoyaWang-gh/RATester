package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func Testlex_1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		left  string
		right string
		want  *lexer
	}{
		{
			name:  "test1",
			input: "test input",
			left:  "{{",
			right: "}}",
			want: &lexer{
				name:         "test1",
				input:        "test input",
				leftDelim:    "{{",
				rightDelim:   "}}",
				line:         1,
				startLine:    1,
				insideAction: false,
			},
		},
		{
			name:  "test2",
			input: "another test",
			left:  "",
			right: "",
			want: &lexer{
				name:         "test2",
				input:        "another test",
				leftDelim:    leftDelim,
				rightDelim:   rightDelim,
				line:         1,
				startLine:    1,
				insideAction: false,
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

			if got := lex(tt.name, tt.input, tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lex() = %v, want %v", got, tt.want)
			}
		})
	}
}
