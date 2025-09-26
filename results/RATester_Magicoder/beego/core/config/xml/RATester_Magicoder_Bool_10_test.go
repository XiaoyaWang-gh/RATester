package xml

import (
	"fmt"
	"testing"
)

func TestBool_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "true",
			"key2": "false",
			"key3": "invalid",
		},
	}

	tests := []struct {
		key      string
		expected bool
		err      error
	}{
		{"key1", true, nil},
		{"key2", false, nil},
		{"key3", false, fmt.Errorf("not exist key: %q", "key3")},
		{"key4", false, fmt.Errorf("not exist key: %q", "key4")},
	}

	for _, test := range tests {
		v, err := cc.Bool(test.key)
		if v != test.expected || err != test.err {
			t.Errorf("expected %v, %v; got %v, %v", test.expected, test.err, v, err)
		}
	}
}
