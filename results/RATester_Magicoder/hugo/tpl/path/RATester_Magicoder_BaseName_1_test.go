package path

import (
	"errors"
	"fmt"
	"testing"
)

func TestBaseName_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name     string
		path     any
		expected string
		err      error
	}{
		{
			name:     "valid path",
			path:     "/path/to/file.txt",
			expected: "file",
			err:      nil,
		},
		{
			name:     "invalid path",
			path:     123,
			expected: "",
			err:      errors.New("unsupported type"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := ns.BaseName(test.path)
			if err != nil && err.Error() != test.err.Error() {
				t.Errorf("expected error %v, got %v", test.err, err)
			}
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
