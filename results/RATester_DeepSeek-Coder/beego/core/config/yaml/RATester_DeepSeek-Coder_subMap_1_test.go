package yaml

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubMap_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		key      string
		expected map[string]interface{}
		err      error
	}

	testCases := []testCase{
		{
			name:     "valid key",
			key:      "key1.key2",
			expected: map[string]interface{}{"key3": "value3", "key4": "value4"},
			err:      nil,
		},
		{
			name:     "invalid key",
			key:      "key5.key6",
			expected: nil,
			err:      fmt.Errorf("the key is invalid: key5.key6"),
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
					"key1": map[string]interface{}{
						"key2": map[string]interface{}{
							"key3": "value3",
							"key4": "value4",
						},
					},
				},
			}

			result, err := cc.subMap(tc.key)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
			if err != nil && tc.err == nil {
				t.Errorf("expected no error, got %v", err)
			}
			if err == nil && tc.err != nil {
				t.Errorf("expected error %v, got no error", tc.err)
			}
			if err != nil && tc.err != nil && err.Error() != tc.err.Error() {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
