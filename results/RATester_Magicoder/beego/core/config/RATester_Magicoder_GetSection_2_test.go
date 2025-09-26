package config

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_2(t *testing.T) {
	fc := &fakeConfigContainer{
		data: map[string]string{
			"section1": "key1=value1,key2=value2",
			"section2": "key3=value3,key4=value4",
		},
	}

	tests := []struct {
		name     string
		section  string
		expected map[string]string
		err      error
	}{
		{
			name:    "TestGetSection_ExistingSection",
			section: "section1",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			err: nil,
		},
		{
			name:     "TestGetSection_NonExistingSection",
			section:  "section3",
			expected: map[string]string{},
			err:      errors.New("section not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := fc.GetSection(test.section)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, but got %v", test.err, err)
			}

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected result %v, but got %v", test.expected, result)
			}
		})
	}
}
