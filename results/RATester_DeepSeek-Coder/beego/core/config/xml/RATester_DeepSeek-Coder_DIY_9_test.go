package xml

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDIY_9(t *testing.T) {
	t.Parallel()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": 123,
			"key3": true,
		},
	}

	tests := []struct {
		name     string
		key      string
		expected interface{}
		err      error
	}{
		{
			name:     "exist key",
			key:      "key1",
			expected: "value1",
			err:      nil,
		},
		{
			name:     "exist key",
			key:      "key2",
			expected: 123,
			err:      nil,
		},
		{
			name:     "exist key",
			key:      "key3",
			expected: true,
			err:      nil,
		},
		{
			name:     "not exist key",
			key:      "key4",
			expected: nil,
			err:      errors.New("not exist key"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			v, err := cc.DIY(tt.key)
			if !reflect.DeepEqual(v, tt.expected) || !reflect.DeepEqual(err, tt.err) {
				t.Errorf("expected %v, got %v", tt.expected, v)
				t.Errorf("expected %v, got %v", tt.err, err)
			}
		})
	}
}
