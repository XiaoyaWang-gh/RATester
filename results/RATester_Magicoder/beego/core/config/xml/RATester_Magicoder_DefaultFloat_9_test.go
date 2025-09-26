package xml

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": float64(1.23),
			"key2": "not a float",
		},
	}

	tests := []struct {
		key        string
		defaultVal float64
		expected   float64
	}{
		{"key1", 0.0, 1.23},
		{"key2", 4.56, 4.56},
		{"key3", 7.89, 7.89},
	}

	for _, test := range tests {
		result := cc.DefaultFloat(test.key, test.defaultVal)
		if result != test.expected {
			t.Errorf("Expected %f, but got %f", test.expected, result)
		}
	}
}
