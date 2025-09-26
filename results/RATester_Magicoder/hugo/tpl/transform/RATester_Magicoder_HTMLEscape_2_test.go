package transform

import (
	"errors"
	"fmt"
	"testing"
)

func TestHTMLEscape_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected string
		err      error
	}{
		{
			name:     "valid input",
			input:    "<script>alert('Hello, world!')</script>",
			expected: "&lt;script&gt;alert('Hello, world!')&lt;/script&gt;",
			err:      nil,
		},
		{
			name:     "invalid input",
			input:    make(chan int),
			expected: "",
			err:      errors.New("unsupported type: chan int"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.HTMLEscape(test.input)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}

			if result != test.expected {
				t.Errorf("Expected result %s, but got %s", test.expected, result)
			}
		})
	}
}
