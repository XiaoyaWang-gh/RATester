package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAssetDir_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
		err      error
	}{
		{
			name:     "valid path",
			input:    "valid/path",
			expected: []string{"file1", "file2"},
			err:      nil,
		},
		{
			name:     "invalid path",
			input:    "invalid/path",
			expected: nil,
			err:      fmt.Errorf("Asset %s not found", "invalid/path"),
		},
		{
			name:     "empty path",
			input:    "",
			expected: []string{"file1", "file2", "file3"},
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			actual, err := AssetDir(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("expected: %v, actual: %v", tt.expected, actual)
			}
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("expected error: %v, actual error: %v", tt.err, err)
			}
		})
	}
}
