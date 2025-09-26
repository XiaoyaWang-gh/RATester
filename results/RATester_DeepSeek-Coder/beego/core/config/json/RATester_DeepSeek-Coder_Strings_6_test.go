package json

import (
	"errors"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
	"go.uber.org/mock/gomock"
)

func TestStrings_6(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		key      string
		expected []string
		err      error
	}{
		{
			name:     "Test case 1",
			key:      "key1",
			expected: []string{"value1", "value2"},
			err:      nil,
		},
		{
			name:     "Test case 2",
			key:      "key2",
			expected: []string{"value3", "value4"},
			err:      nil,
		},
		{
			name:     "Test case 3",
			key:      "key3",
			expected: nil,
			err:      errors.New("key not found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			c := &JSONConfigContainer{
				data: map[string]interface{}{
					tc.key: tc.expected,
				},
			}

			result, err := c.Strings(tc.key)
			if tc.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.err, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
