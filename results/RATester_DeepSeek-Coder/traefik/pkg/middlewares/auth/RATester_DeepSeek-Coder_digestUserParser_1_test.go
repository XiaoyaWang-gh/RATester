package auth

import (
	"fmt"
	"testing"
)

func TestDigestUserParser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type test struct {
		input    string
		expected string
		err      error
	}

	tests := []test{
		{"user:realm:nonce", "user:realm", nil},
		{"user:realm:nonce:extra", "", nil},
		{"user", "", nil},
	}

	for _, test := range tests {
		result, _, err := digestUserParser(test.input)
		if err != nil && test.err == nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if err == nil && test.err != nil {
			t.Errorf("Expected error, got none")
		}
		if result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}
