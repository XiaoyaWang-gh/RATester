package json

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_6(t *testing.T) {
	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"section1": map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			"section2": map[string]string{
				"key3": "value3",
				"key4": "value4",
			},
		},
	}

	tests := []struct {
		name     string
		section  string
		expected map[string]string
		err      error
	}{
		{
			name:    "ExistingSection",
			section: "section1",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			err: nil,
		},
		{
			name:    "NonExistingSection",
			section: "nonexist",
			err:     errors.New("nonexist section nonexist"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := container.GetSection(test.section)
			if err != nil {
				if err.Error() != test.err.Error() {
					t.Errorf("Expected error %v, but got %v", test.err, err)
				}
			} else {
				if !reflect.DeepEqual(result, test.expected) {
					t.Errorf("Expected result %v, but got %v", test.expected, result)
				}
			}
		})
	}
}
