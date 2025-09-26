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
			name:     "valid directory",
			input:    "valid/directory",
			expected: []string{"file1", "file2"},
			err:      nil,
		},
		{
			name:     "invalid directory",
			input:    "invalid/directory",
			expected: nil,
			err:      fmt.Errorf("Asset %s not found", "invalid/directory"),
		},
		{
			name:     "empty directory",
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
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
			if err != nil && tt.err == nil {
				t.Errorf("expected no error, got %v", err)
			}
			if err == nil && tt.err != nil {
				t.Errorf("expected error %v, got no error", tt.err)
			}
			if err != nil && tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("expected error %v, got %v", tt.err, err)
			}
		})
	}
}
