package yaml

import (
	"errors"
	"fmt"
	"testing"
)

func TestInt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": 123,
			"key2": int64(456),
			"key3": "not int",
		},
	}

	tests := []struct {
		key      string
		expected int
		err      error
	}{
		{"key1", 123, nil},
		{"key2", 456, nil},
		{"key3", 0, errors.New("not int value")},
		{"key4", 0, errors.New("not int value")},
	}

	for _, test := range tests {
		got, err := cc.Int(test.key)
		if err != test.err {
			t.Errorf("Int(%q) got error %v, want %v", test.key, err, test.err)
		}
		if got != test.expected {
			t.Errorf("Int(%q) got %v, want %v", test.key, got, test.expected)
		}
	}
}
