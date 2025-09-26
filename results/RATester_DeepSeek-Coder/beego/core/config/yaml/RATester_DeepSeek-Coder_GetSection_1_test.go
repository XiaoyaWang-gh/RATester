package yaml

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		section  string
		expected map[string]string
		err      error
	}

	testCases := []testCase{
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
			name:    "non-existing section",
			section: "section3",
			expected: map[string]string{
				"key3": "value3",
			},
			err: errors.New("not exist section"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cc := &ConfigContainer{
				data: map[string]interface{}{
					"section1": map[string]string{
						"key1": "value1",
						"key2": "value2",
					},
					"section2": map[string]interface{}{
						"key3": 3,
						"key4": 4.0,
					},
				},
			}

			res, err := cc.GetSection(tc.section)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, res)
			}
			if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
