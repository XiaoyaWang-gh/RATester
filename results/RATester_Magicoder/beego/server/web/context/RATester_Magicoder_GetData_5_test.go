package context

import (
	"fmt"
	"testing"
)

func TestGetData_5(t *testing.T) {
	input := &BeegoInput{
		data: map[interface{}]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	tests := []struct {
		name  string
		key   interface{}
		value interface{}
	}{
		{"TestGetData_1", "key1", "value1"},
		{"TestGetData_2", "key2", "value2"},
		{"TestGetData_3", "key3", nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := input.GetData(test.key)
			if result != test.value {
				t.Errorf("Expected %v, got %v", test.value, result)
			}
		})
	}
}
