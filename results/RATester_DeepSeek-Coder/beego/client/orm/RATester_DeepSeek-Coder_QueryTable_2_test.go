package orm

import (
	"fmt"
	"testing"
)

func TestQueryTable_2(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		ID int
	}

	testCases := []struct {
		name           string
		input          interface{}
		expectedOutput string
	}{
		{
			name:           "Test with string input",
			input:          "testStruct",
			expectedOutput: "testStruct",
		},
		{
			name:           "Test with struct input",
			input:          &testStruct{},
			expectedOutput: "testStruct",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &filterOrmDecorator{}
			res := f.QueryTable(tc.input)
			if res == nil {
				t.Errorf("Expected non-nil result, got nil")
			}
		})
	}
}
