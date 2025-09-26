package parse

import (
	"fmt"
	"testing"
)

func TestnextItem_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected item
	}{
		{
			name:  "EOF",
			input: "",
			expected: item{
				typ:  itemEOF,
				pos:  0,
				val:  "EOF",
				line: 1,
			},
		},
		{
			name:  "Text",
			input: "Hello, World!",
			expected: item{
				typ:  itemText,
				pos:  0,
				val:  "Hello, World!",
				line: 1,
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

			l := &lexer{
				name:         "test",
				input:        test.input,
				start:        0,
				line:         1,
				item:         item{},
				pos:          0,
				atEOF:        false,
				insideAction: false,
			}

			result := l.nextItem()

			if result.typ != test.expected.typ {
				t.Errorf("Expected item type %d, but got %d", test.expected.typ, result.typ)
			}

			if result.pos != test.expected.pos {
				t.Errorf("Expected item position %d, but got %d", test.expected.pos, result.pos)
			}

			if result.val != test.expected.val {
				t.Errorf("Expected item value %s, but got %s", test.expected.val, result.val)
			}

			if result.line != test.expected.line {
				t.Errorf("Expected item line %d, but got %d", test.expected.line, result.line)
			}
		})
	}
}
