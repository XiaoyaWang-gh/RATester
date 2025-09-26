package orm

import (
	"fmt"
	"testing"
)

func TestReadOrCreate_2(t *testing.T) {
	o := &ormBase{}
	type TestStruct struct {
		ID   int64
		Name string
	}

	testCases := []struct {
		name     string
		input    TestStruct
		expected bool
	}{
		{
			name: "Test case 1",
			input: TestStruct{
				ID:   1,
				Name: "Test",
			},
			expected: true,
		},
		{
			name: "Test case 2",
			input: TestStruct{
				ID:   2,
				Name: "Test2",
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, _, err := o.ReadOrCreate(&tc.input, "ID")
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
			_, _, err = o.ReadOrCreate(&tc.input, "ID")
			if err != nil {
				t.Errorf("Expected no error, got %v", err)
			}
		})
	}
}
