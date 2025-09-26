package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateValueFuncs_1(t *testing.T) {
	testCases := []struct {
		name     string
		funcMap  FuncMap
		expected map[string]reflect.Value
	}{
		{
			name: "Test case 1",
			funcMap: FuncMap{
				"func1": func() {},
				"func2": func() error { return nil },
			},
			expected: map[string]reflect.Value{
				"func1": reflect.ValueOf(func() {}),
				"func2": reflect.ValueOf(func() error { return nil }),
			},
		},
		{
			name: "Test case 2",
			funcMap: FuncMap{
				"func3": func(int) {},
				"func4": func() (int, error) { return 0, nil },
			},
			expected: map[string]reflect.Value{
				"func3": reflect.ValueOf(func(int) {}),
				"func4": reflect.ValueOf(func() (int, error) { return 0, nil }),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := createValueFuncs(tc.funcMap)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
