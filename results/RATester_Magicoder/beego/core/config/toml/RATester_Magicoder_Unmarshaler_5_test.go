package toml

import (
	"errors"
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestUnmarshaler_5(t *testing.T) {
	type testStruct struct {
		Field1 string
		Field2 int
	}

	testCases := []struct {
		name     string
		prefix   string
		obj      interface{}
		expected error
	}{
		{
			name:     "Test case 1",
			prefix:   "test.prefix",
			obj:      &testStruct{},
			expected: nil,
		},
		{
			name:     "Test case 2",
			prefix:   "",
			obj:      &testStruct{},
			expected: nil,
		},
		{
			name:     "Test case 3",
			prefix:   "invalid.prefix",
			obj:      &testStruct{},
			expected: errors.New("invalid prefix"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			cc := &configContainer{
				t: &toml.Tree{},
			}

			err := cc.Unmarshaler(tc.prefix, tc.obj)

			if err != nil && tc.expected == nil {
				t.Errorf("Expected no error, but got: %v", err)
			}

			if err == nil && tc.expected != nil {
				t.Errorf("Expected error: %v, but got no error", tc.expected)
			}

			if err != nil && tc.expected != nil && err.Error() != tc.expected.Error() {
				t.Errorf("Expected error: %v, but got: %v", tc.expected, err)
			}
		})
	}
}
