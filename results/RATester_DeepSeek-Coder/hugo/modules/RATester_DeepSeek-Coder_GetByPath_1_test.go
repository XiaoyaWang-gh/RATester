package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetByPath_1(t *testing.T) {
	type test struct {
		name     string
		modules  goModules
		path     string
		expected *goModule
	}

	tests := []test{
		{
			name: "Test with existing path",
			modules: goModules{
				{Path: "github.com/test/test1", Version: "v1.0.0"},
				{Path: "github.com/test/test2", Version: "v2.0.0"},
			},
			path:     "github.com/test/test1",
			expected: &goModule{Path: "github.com/test/test1", Version: "v1.0.0"},
		},
		{
			name: "Test with non-existing path",
			modules: goModules{
				{Path: "github.com/test/test1", Version: "v1.0.0"},
				{Path: "github.com/test/test2", Version: "v2.0.0"},
			},
			path:     "github.com/test/test3",
			expected: nil,
		},
		{
			name: "Test with case-insensitive path",
			modules: goModules{
				{Path: "github.com/test/test1", Version: "v1.0.0"},
				{Path: "github.com/test/test2", Version: "v2.0.0"},
			},
			path:     "GiTHuB.CoM/TeSt/TeSt1",
			expected: &goModule{Path: "github.com/test/test1", Version: "v1.0.0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.modules.GetByPath(tt.path)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}
