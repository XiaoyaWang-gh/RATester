package parse

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLex_1(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		left   string
		right  string
		expect *lexer
	}{
		{
			name:  "test1",
			input: "test input",
			left:  "{{",
			right: "}}",
			expect: &lexer{
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
			expect: &lexer{
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := lex(test.name, test.input, test.left, test.right)
			if !reflect.DeepEqual(got, test.expect) {
				t.Errorf("got %v, want %v", got, test.expect)
			}
		})
	}
}
