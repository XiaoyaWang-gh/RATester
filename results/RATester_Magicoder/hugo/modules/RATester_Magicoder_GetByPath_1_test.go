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
			name: "TestGetByPath_ExistingModule",
			modules: goModules{
				&goModule{Path: "github.com/test/module1"},
				&goModule{Path: "github.com/test/module2"},
			},
			path:     "github.com/test/module1",
			expected: &goModule{Path: "github.com/test/module1"},
		},
		{
			name: "TestGetByPath_NonExistingModule",
			modules: goModules{
				&goModule{Path: "github.com/test/module1"},
				&goModule{Path: "github.com/test/module2"},
			},
			path:     "github.com/test/module3",
			expected: nil,
		},
		{
			name:     "TestGetByPath_NilModules",
			modules:  nil,
			path:     "github.com/test/module1",
			expected: nil,
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
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
