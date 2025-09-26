package json

import (
	"errors"
	"fmt"
	"testing"
)

func TestInt64_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": float64(123),
			"key2": "not a number",
			"key3": nil,
		},
	}

	tests := []struct {
		key      string
		expected int64
		err      error
	}{
		{"key1", 123, nil},
		{"key2", 0, errors.New("not int64 value")},
		{"key3", 0, errors.New("not exist key:key3")},
		{"key4", 0, errors.New("not exist key:key4")},
	}

	for _, test := range tests {
		val, err := container.Int64(test.key)
		if val != test.expected || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf("Expected (%v, %v), got (%v, %v)", test.expected, test.err, val, err)
		}
	}
}
