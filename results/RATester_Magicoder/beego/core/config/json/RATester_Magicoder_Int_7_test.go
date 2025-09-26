package json

import (
	"errors"
	"fmt"
	"testing"
)

func TestInt_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": 123,
			"key2": "456",
			"key3": 789.0,
			"key4": "abc",
		},
	}

	tests := []struct {
		key      string
		expected int
		err      error
	}{
		{"key1", 123, nil},
		{"key2", 456, nil},
		{"key3", 789, nil},
		{"key4", 0, errors.New("not valid value")},
		{"key5", 0, errors.New("not exist key:key5")},
	}

	for _, test := range tests {
		val, err := container.Int(test.key)
		if val != test.expected || err != test.err {
			t.Errorf("Expected (%d, %v), got (%d, %v)", test.expected, test.err, val, err)
		}
	}
}
