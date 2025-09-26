package yaml

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestSub_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		data     map[string]interface{}
		key      string
		expected map[string]interface{}
		err      error
	}{
		{
			name: "success",
			data: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			key: "key1",
			expected: map[string]interface{}{
				"key2": "value2",
			},
			err: nil,
		},
		{
			name: "key not exist",
			data: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			key:      "key3",
			expected: nil,
			err:      errors.New("key not exist"),
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
			sub, err := c.Sub(tc.key)
			if tc.err != nil {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expected, sub.(*ConfigContainer).data)
		})
	}
}
