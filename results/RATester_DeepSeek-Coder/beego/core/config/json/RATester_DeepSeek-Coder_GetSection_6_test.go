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
			name:    "existing section",
			section: "section1",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			err: nil,
		},
		{
			name:     "nonexisting section",
			section:  "section3",
			expected: nil,
			err:      errors.New("nonexist section section3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := container.GetSection(tt.section)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}
