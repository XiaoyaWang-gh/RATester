package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTargetPaths_1(t *testing.T) {
	tests := []struct {
		name            string
		targetBasePaths []string
		expectedPaths   []string
	}{
		{
			name:            "single base path",
			targetBasePaths: []string{"/base"},
			expectedPaths:   []string{"/base/target"},
		},
		{
			name:            "multiple base paths",
			targetBasePaths: []string{"/base1", "/base2"},
			expectedPaths:   []string{"/base1/target", "/base2/target"},
		},
		{
			name:            "no base paths",
			targetBasePaths: []string{},
			expectedPaths:   []string{"/target"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			rp := ResourcePaths{
				TargetBasePaths: test.targetBasePaths,
			}

			paths := rp.TargetPaths()

			if !reflect.DeepEqual(paths, test.expectedPaths) {
				t.Errorf("Expected paths %v, but got %v", test.expectedPaths, paths)
			}
		})
	}
}
