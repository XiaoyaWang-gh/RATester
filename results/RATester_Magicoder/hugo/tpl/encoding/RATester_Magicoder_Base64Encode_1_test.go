package encoding

import (
	"errors"
	"fmt"
	"testing"
)

func TestBase64Encode_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		input    any
		expected string
		err      error
	}{
		{
			name:     "valid input",
			input:    "test",
			expected: "dGVzdA==",
			err:      nil,
		},
		{
			name:     "invalid input",
			input:    123,
			expected: "",
			err:      errors.New("unsupported type"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.Base64Encode(test.input)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("expected error %v, got %v", test.err, err)
			}

			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
