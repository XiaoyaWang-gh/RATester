package yaml

import (
	"fmt"
	"testing"
)

func TestDefaultString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": "",
			"key3": "value3",
		},
	}

	tests := []struct {
		key        string
		defaultVal string
		expected   string
	}{
		{"key1", "default1", "value1"},
		{"key2", "default2", "default2"},
		{"key3", "default3", "value3"},
		{"key4", "default4", "default4"},
	}

	for _, test := range tests {
		result := cc.DefaultString(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("DefaultString(%s, %s) = %s; want %s", test.key, test.defaultVal, result, test.expected)
		}
	}
}
