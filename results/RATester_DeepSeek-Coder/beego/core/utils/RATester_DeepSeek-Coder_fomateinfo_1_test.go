package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestFomateinfo_1(t *testing.T) {
	type testStruct struct {
		Name string
		Age  int
	}

	testCases := []struct {
		name     string
		headlen  int
		data     []interface{}
		expected []byte
	}{
		{
			name:    "single value",
			headlen: 4,
			data:    []interface{}{"test"},
			expected: []byte(`    [
        "test"
    ]`),
		},
		{
			name:    "multiple values",
			headlen: 4,
			data:    []interface{}{"test1", "test2"},
			expected: []byte(`    [
        "test1", 
        "test2"
    ]`),
		},
		{
			name:    "struct value",
			headlen: 4,
			data:    []interface{}{testStruct{"John", 30}},
			expected: []byte(`    [
        {
            "Name": "John",
            "Age": 30
        }
    ]`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := fomateinfo(tc.headlen, tc.data...)
			if !bytes.Equal(result, tc.expected) {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
