package dartsass

import (
	"fmt"
	"testing"
)

func TestKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testCases := []struct {
		name     string
		optsm    map[string]any
		expected string
	}{
		{
			name:     "test1",
			optsm:    map[string]any{"key1": "value1", "key2": "value2"},
			expected: "test1:value1:value2",
		},
		{
			name:     "test2",
			optsm:    map[string]any{"key3": "value3", "key4": "value4"},
			expected: "test2:value3:value4",
		},
	}

	for _, tc := range testCases {
		transform := &transform{
			optsm: tc.optsm,
		}

		key := transform.Key()
		if key.Value() != tc.expected {
			t.Errorf("Expected %s, but got %s", tc.expected, key.Value())
		}
	}
}
