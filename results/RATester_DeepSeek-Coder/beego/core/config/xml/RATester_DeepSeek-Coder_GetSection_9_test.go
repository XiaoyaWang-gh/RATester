package xml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_9(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		data     map[string]interface{}
		section  string
		expected map[string]string
		err      error
	}

	testCases := []testCase{
		{
			name: "success",
			data: map[string]interface{}{
				"section1": map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			section: "section1",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			err: nil,
		},
		{
			name: "section not found",
			data: map[string]interface{}{
				"section1": map[string]interface{}{
					"key1": "value1",
					"key2": "value2",
				},
			},
			section:  "section2",
			expected: nil,
			err:      fmt.Errorf("section 'section2' not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ConfigContainer{
				data: tc.data,
			}
			result, err := c.GetSection(tc.section)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
			if !reflect.DeepEqual(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
